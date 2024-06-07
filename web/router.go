package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/iypetrov/goshop/internal/common"
	"net/http"
)

func Router() http.Handler {
	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Get("/home", HomeHandler)
	r.Get("/register", RegisterHandler)
	r.Get("/login", LoginHandler)
	r.Get("/404", NotFoundHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		common.RedirectNotFoundView(w, r)
	})

	r.Get("/user/{id}", HomeHandler)

	return r
}
