package users

import (
	"context"
	"errors"
	"fmt"
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/iypetrov/goshop/internal/common"
	"github.com/iypetrov/goshop/internal/config"
	"github.com/iypetrov/goshop/web/templates/views"
	"github.com/markbates/goth/gothic"
	"net/http"
)

func createAuthToken(w http.ResponseWriter, r *http.Request, response map[string]interface{}) error {
	_, token, err := config.Get().Auth.TokenAuth.Encode(response)
	if err != nil {
		return err
	}

	session, err := config.Get().Auth.Store.Get(r, common.CookieName)
	if err != nil {
		return err
	}
	session.Values["token"] = token

	if err := session.Save(r, w); err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	return nil
}

func Provider(w http.ResponseWriter, r *http.Request) error {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	if _, err := gothic.CompleteUserAuth(w, r); err != nil {
		gothic.BeginAuthHandler(w, r)
	} else {
		templ.Handler(views.Login(config.Get().GetBaseAPIURL()))
	}

	return nil
}

func ProviderCallback(w http.ResponseWriter, r *http.Request) error {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return FailedAuthUser()
	}

	user, err := GetServer().Auth(gothUser.Email, provider)
	if err != nil {
		return FailedAuthUser()
	}

	response := CreateResponseDTOFromModel(user)
	if err = createAuthToken(w, r, response.ToString()); err != nil {
		return FailedGenerateAuthToken()
	}

	common.RedirectHomeView(w, r)

	return nil
}

func GetByIDHandler(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")

	model, err := GetServer().GetModelByID(id)
	if err != nil {
		return InvalidID()
	}

	return common.WriteJSON(w, http.StatusOK, CreateResponseDTOFromModel(model))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) error {
	var requestDTO RegisterRequestDTO
	cb, err := common.ReadRequestBody(r, &requestDTO)
	defer cb()
	if err != nil {
		return err
	}

	model, err := GetServer().CreateModel(CreateModelFromRegisterRequestDTO(requestDTO))
	if err != nil {
		if errors.As(err, &common.GeneralError{}) {
			return common.InvalidRequestData(err)
		}
		return FailedCreation()
	}

	response := CreateResponseDTOFromModel(model)
	if err = createAuthToken(w, r, response.ToString()); err != nil {
		return FailedGenerateAuthToken()
	}

	return common.WriteJSON(w, http.StatusOK, response)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) error {
	var requestDTO LoginRequestDTO
	cb, err := common.ReadRequestBody(r, &requestDTO)
	defer cb()
	if err != nil {
		return err
	}

	model, err := GetServer().FindModelByEmailAndAuthProvider(CreateModelFromLoginRequestDTO(requestDTO))
	if err != nil {
		if errors.As(err, &common.GeneralError{}) {
			return common.InvalidRequestData(err)
		}
		return FailedFind()
	}

	response := CreateResponseDTOFromModel(model)
	if err = createAuthToken(w, r, response.ToString()); err != nil {
		return FailedGenerateAuthToken()
	}

	return common.WriteJSON(w, http.StatusOK, CreateResponseDTOFromModel(model))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) error {
	err := gothic.Logout(w, r)
	if err != nil {
		return FailedLogout()
	}
	r = r.WithContext(context.WithValue(r.Context(), "provider", ""))

	common.RedirectLoginView(w, r)

	return nil
}
