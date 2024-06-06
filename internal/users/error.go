package users

import (
	"fmt"
	"github.com/iypetrov/goshop/internal/common"
	"net/http"
)

func FailedAuthUser() common.APIError {
	return common.APIError{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprint("Failed to authenticate a user"),
	}
}

func InvalidID() common.APIError {
	return common.APIError{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprint("User with this ID doesn't exist"),
	}
}

func FailedCreation() common.APIError {
	return common.APIError{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprint("Failed to create a user"),
	}
}

func FailedFind() common.APIError {
	return common.APIError{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprint("User doesn't exist"),
	}
}

func FailedLogout() common.APIError {
	return common.APIError{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprint("Failed to logout"),
	}
}
