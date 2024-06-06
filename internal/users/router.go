package users

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	ch "github.com/iypetrov/goshop/internal/common"
)

func Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", ch.MakeHandler(GetByIDHandler))
	return r
}

func AuthRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/{provider}", ch.MakeHandler(Provider))
	r.Get("/{provider}/callback", ch.MakeHandler(ProviderCallback))
	return r
}

func OnboardingRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/register", ch.MakeHandler(RegisterHandler))
	r.Post("/login", ch.MakeHandler(LoginHandler))
	r.Get("/logout", ch.MakeHandler(LogoutHandler))
	return r
}

func AdminRouter() http.Handler {
	r := chi.NewRouter()
	return r
}
