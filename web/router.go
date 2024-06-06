package web

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Router() http.Handler {
	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Get("/home/{id}", HomeHandler)
	r.Get("/register", RegisterHandler)
	r.Get("/login", LoginHandler)
	return r
}
