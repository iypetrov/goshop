package web

import (
	"github.com/a-h/templ"
	"github.com/iypetrov/goshop/internal/common"
	"github.com/iypetrov/goshop/internal/config"
	"github.com/iypetrov/goshop/web/templates/views"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.Home()).ServeHTTP(w, r)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.Register(config.Get().GetBaseAPIURL())).ServeHTTP(w, r)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.Login(config.Get().GetBaseAPIURL())).ServeHTTP(w, r)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	token, err := common.GetToken(r, config.Get().Auth.Store, common.CookieName)
	if err != nil {
		common.RedirectLoginView(w, r)
		return
	}

	data, err := config.Get().Auth.TokenAuth.Decode(token)
	if err != nil {
		common.RedirectLoginView(w, r)
		return
	}

	id, ok := data.Get("id")
	if !ok {
		common.RedirectLoginView(w, r)
		return
	}

	templ.Handler(views.NotFound(config.Get().GetBaseWebURL(), id.(string))).ServeHTTP(w, r)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	templ.Handler(views.Home(id)).ServeHTTP(w, r)
}
