package util

import (
	"net/http"

	"gotiny/internal/api/dto"
	"gotiny/internal/core/model"
	"gotiny/internal/core/usecase"
)

var statusCodeMapping = map[string]int{
	model.UnauthorizedError: http.StatusUnauthorized,
	model.NotFoundError:     http.StatusNotFound,
	model.UnknownError:      http.StatusInternalServerError,
	model.InvalidInputError: http.StatusBadRequest,
	usecase.InvalidUrlError: http.StatusBadRequest,
}

var errorMessageMapping = map[string]string{
	usecase.InvalidUrlError: "Invalid url",
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

func WriteAjaxError(writer http.ResponseWriter, err error) {
	var message string
	switch err.(type) {
	case *model.AppError:
		appError := err.(*model.AppError)
		m, ok := errorMessageMapping[appError.Type]
		if ok {
			message = m
		}
	}
	if message == "" {
		message = "Something went wrong"
	}
	statusCode := statusCodeFromError(err)
	writer.WriteHeader(statusCode)
	writer.Write([]byte(message))
}
