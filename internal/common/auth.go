package common

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

const (
	CookieName = "jwt"
)

func GetToken(r *http.Request, store *sessions.CookieStore, cookieName string) (string, error) {
	session, err := store.Get(r, cookieName)
	if err != nil {
		return "", fmt.Errorf("no cookies attached to the request")
	}

	tokenValue, ok := session.Values["token"]
	if !ok {
		return "", fmt.Errorf("token is not found")
	}

	token, ok := tokenValue.(string)
	if !ok {
		return "", fmt.Errorf("token is not in correct state")
	}

	return token, nil
}
