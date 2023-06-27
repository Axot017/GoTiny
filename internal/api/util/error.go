package util

import (
	"net/http"

	"gotiny/internal/api/dto"
	"gotiny/internal/core/model"
)

var statusCodeMapping = map[string]int{
	model.UnauthorizedError: http.StatusUnauthorized,
	model.NotFoundError:     http.StatusNotFound,
	model.UnknownError:      http.StatusInternalServerError,
	model.InvalidInputError: http.StatusBadRequest,
}

func statusCodeFromError(err error) int {
	switch err.(type) {
	case *model.AppError:
		appError := err.(*model.AppError)
		code, ok := statusCodeMapping[appError.Type]
		if !ok {
			code = 500
		}
		return code
	default:
		return 500
	}
}

func WriteError(writer http.ResponseWriter, err error) {
	dto := dto.ErrorDtoFromError(err)
	statusCode := statusCodeFromError(err)
	WriteResponseJson(writer, dto, statusCode)
}
