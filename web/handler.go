package web

import (
	"github.com/a-h/templ"
	"github.com/iypetrov/goshop/internal/config"
	"github.com/iypetrov/goshop/web/templates/views"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	templ.Handler(views.Home(id)).ServeHTTP(w, r)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.Register(config.Get().GetBaseURL())).ServeHTTP(w, r)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.Login(config.Get().GetBaseURL())).ServeHTTP(w, r)
}
