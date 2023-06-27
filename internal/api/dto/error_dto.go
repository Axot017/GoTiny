package dto

import "gotiny/internal/core/model"

type ErrorDto struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Args    map[string]string `json:"args,omitempty"`
}

var typeMapping map[string]string = map[string]string{
	model.UnknownError:      "internal_server_error",
	model.NotFoundError:     "not_found",
	model.UnauthorizedError: "unauthorized",
	model.InvalidInputError: "input_error",
}

var messageMapping map[string]string = map[string]string{
	model.UnknownError:      "Internal server error",
	model.NotFoundError:     "Not found",
	model.UnauthorizedError: "Unauthorized",
	model.InvalidInputError: "Invalid input",
}

func ErrorDtoFromError(err error) ErrorDto {
	switch err.(type) {
	case *model.AppError:
		appError := err.(*model.AppError)
		code, ok := typeMapping[appError.Type]
		if !ok {
			code = "unknown_server_error"
		}
		message, ok := messageMapping[appError.Type]
		if !ok {
			message = "Unknown server error"
		}
		return ErrorDto{
			Code:    code,
			Message: message,
			Args:    appError.Args,
		}
	default:
		return ErrorDto{
			Code:    "unknown_server_error",
			Message: "Unknown server error",
		}
	}
}
