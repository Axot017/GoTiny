package model

import "fmt"

const (
	UnknownError      string = "unknown_error"
	NotFoundError     string = "not_found_error"
	UnauthorizedError string = "unauthorized_error"
	InvalidInputError string = "invalid_input_error"
)

type AppError struct {
	Type    string
	Context error
	Args    map[string]string
}

func (e *AppError) Error() string {
	if e.Context == nil {
		return fmt.Sprintf("%s", e.Type)
	}

	return fmt.Sprintf("%s: %s", e.Type, e.Context.Error())
}

func NewNotFoundError() *AppError {
	return &AppError{
		Type: NotFoundError,
	}
}
