package common

import (
	"fmt"
	"github.com/iypetrov/goshop/internal/config"
	"net/http"
)

func RedirectHomeView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "/")
	http.Redirect(w, r, fmt.Sprintf("%s/home", config.Get().GetBaseWebURL()), http.StatusTemporaryRedirect)
}

func RedirectLoginView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "/")
	http.Redirect(w, r, fmt.Sprintf("%s/login", config.Get().GetBaseWebURL()), http.StatusTemporaryRedirect)
}

func RedirectNotFoundView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "/")
	http.Redirect(w, r, fmt.Sprintf("%s/404", config.Get().GetBaseWebURL()), http.StatusTemporaryRedirect)
}
