package users

import (
	"github.com/iypetrov/goshop/internal/common"
	"github.com/iypetrov/goshop/internal/config"
	"net/http"
)

func PrivateAuthenticator() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			token, err := common.GetToken(r, config.Get().Auth.Store, common.CookieName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}

			_, err = config.Get().Auth.TokenAuth.Decode(token)
			if err != nil {
				http.Error(w, "Not valid token", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(hfn)
	}
}

func AdminAuthenticator() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			token, err := common.GetToken(r, config.Get().Auth.Store, common.CookieName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}

			data, err := config.Get().Auth.TokenAuth.Decode(token)
			if err != nil {
				http.Error(w, "Not valid token", http.StatusUnauthorized)
				return
			}

			role, ok := data.Get("user_role")
			if !ok {
				http.Error(w, "Token is in not correct state", http.StatusUnauthorized)
				return
			}

			if role.(string) == ADMIN {
				http.Error(w, "Only admins have access", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(hfn)
	}
}
