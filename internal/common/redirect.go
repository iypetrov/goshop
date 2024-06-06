package common

import (
	"fmt"
	"github.com/iypetrov/goshop/internal/config"
	"net/http"
)

func RedirectHomeView(w http.ResponseWriter, r *http.Request, userID string) {
	w.Header().Set("Location", "/")
	http.Redirect(w, r, fmt.Sprintf("%s/home/%s", config.Get().GetBaseWebURL(), userID), http.StatusTemporaryRedirect)
}

func RedirectLoginView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "/")
	http.Redirect(w, r, fmt.Sprintf("%s/login", config.Get().GetBaseWebURL()), http.StatusTemporaryRedirect)
}
