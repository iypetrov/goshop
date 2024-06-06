package common

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
)

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func MakeHandler(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			var apiErr APIError
			if errors.As(err, &apiErr) {
				WriteJSON(w, apiErr.StatusCode, apiErr)
			}
			slog.Error(err.Error())
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func ReadRequestBody(r *http.Request, request interface{}) (func(), error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, FailedReadRequestBody()
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		return nil, InvalidJSON()
	}

	return func() {
		r.Body.Close()
	}, nil
}
