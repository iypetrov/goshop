package common

import (
	"fmt"
	"net/http"
)

type GeneralError struct {
	Message any `json:"message"`
}

func (e GeneralError) Error() string {
	return fmt.Sprintf("General error: %d", e.Message)
}

func FailedValidation(errors error) GeneralError {
	return GeneralError{
		Message: errors.Error(),
	}
}

type APIError struct {
	StatusCode int `json:"statusCode"`
	Message    any `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("API error: %d", e.StatusCode)
}

func InvalidRequestData(errors error) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    errors.Error(),
	}
}

func InvalidJSON() APIError {
	return APIError{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprint("Invalid JSON request data"),
	}
}

func FailedReadRequestBody() APIError {
	return APIError{
		StatusCode: http.StatusInternalServerError,
		Message:    fmt.Sprint("Failed to read a request body"),
	}
}
