package users

import (
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/iypetrov/goshop/internal/common"
	"github.com/iypetrov/goshop/internal/config"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"sync"
)

type Server struct {
	tokenAuth  *jwtauth.JWTAuth
	store      *sessions.CookieStore
	repository *Repository
}

var (
	server  *Server
	oncesrv sync.Once
)

func NewServer(repository *Repository) {
	oncesrv.Do(func() {
		server = &Server{
			tokenAuth:  jwtauth.New("HS256", []byte(config.Get().Auth.AuthKey), nil),
			store:      sessions.NewCookieStore([]byte(config.Get().Auth.AuthKey)),
			repository: repository,
		}
	})
}

func GetServer() *Server {
	return server
}

func GetTokenAuth() *jwtauth.JWTAuth {
	return server.tokenAuth
}

func GetStore() *sessions.CookieStore {
	return server.store
}

func (s *Server) InitAuthProviders() {
	server.store.MaxAge(86400)
	server.store.Options.Path = "/"
	server.store.Options.HttpOnly = true
	server.store.Options.Secure = config.Get().App.Environment != "dev"
	gothic.Store = server.store

	goth.UseProviders(
		google.New(
			config.Get().Auth.GoogleClientID,
			config.Get().Auth.GoogleClientSecret,
			fmt.Sprintf("%s/auth/google/callback", config.Get().GetBaseAPIURL()),
		),
		github.New(
			config.Get().Auth.GithubClientID,
			config.Get().Auth.GithubClientSecret,
			fmt.Sprintf("%s/auth/github/callback", config.Get().GetBaseAPIURL()),
		),
		facebook.New(
			config.Get().Auth.FacebookClientID,
			config.Get().Auth.FacebookClientSecret,
			fmt.Sprintf("%s/auth/facebook/callback", config.Get().GetBaseAPIURL()),
		),
	)
}

func (s *Server) Auth(email string, provider string) (Model, error) {
	provider = strings.ToUpper(provider)

	loginModel := CreateModelFromLoginRequestDTO(LoginRequestDTO{
		Email:        email,
		Password:     "",
		AuthProvider: ConvertToAuthProvider(provider),
	})
	registerModel := CreateModelFromRegisterRequestDTO(RegisterRequestDTO{
		Email:        email,
		Password:     "",
		AuthProvider: ConvertToAuthProvider(provider),
	})

	user, err := GetServer().FindModelByEmailAndAuthProvider(loginModel)
	if err != nil {
		user, err = GetServer().CreateModel(registerModel)
		if err != nil {
			return Model{}, err
		}
	}

	return user, nil
}

func (s *Server) GetModelByID(id string) (Model, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return Model{}, err
	}

	entity, err := s.repository.GetEntityByID(uuidID)
	if err != nil {
		return Model{}, err
	}

	return CreateModelFromEntity(entity), nil
}

func (s *Server) CreateModel(model Model) (Model, error) {
	if err := model.Validate(); err != nil {
		return Model{}, err
	}

	if model.AuthProvider == NONE {
		password, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
		if err != nil {
			return Model{}, err
		}
		model.Password = string(password)
	} else {
		// We expect the received password to be empty since the user uses an auth provider, but to ensure no unhashed passwords are stored in the database.
		model.Password = ""
	}

	entity, err := s.repository.CreateEntity(CreateEntityFromModel(model))
	if err != nil {
		return Model{}, err
	}

	return CreateModelFromEntity(entity), nil
}

func (s *Server) FindModelByEmailAndAuthProvider(model Model) (Model, error) {
	if err := model.Validate(); err != nil {
		return Model{}, common.FailedValidation(err)
	}

	entity, err := s.repository.GetEntityByEmail(model.Email, model.AuthProvider)
	if err != nil {
		return Model{}, err
	}

	if model.AuthProvider == NONE {
		if bcrypt.CompareHashAndPassword([]byte(entity.Password), []byte(model.Password)) != nil {
			return Model{}, err
		}
	}

	return CreateModelFromEntity(entity), nil
}
