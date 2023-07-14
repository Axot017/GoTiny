package util

import (
	"net/http"

	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator/v10"

	"gotiny/internal/core/model"
)

func DecodeAndValidateFrom[T any](
	r *http.Request,
	validate *validator.Validate,
	decoder *form.Decoder,
) (T, error) {
	dto, err := DecodeForm[T](r, decoder)
	if err != nil {
		return dto, err
	}

	err = validate.Struct(dto)

	if err != nil {
		return dto, getInvalidInputError(err)
	}
	return dto, nil
}

func DecodeForm[T any](r *http.Request, decoder *form.Decoder) (T, error) {
	var dto T
	err := r.ParseForm()
	if err != nil {
		return dto, getInvalidInputError(err)
	}
	err = decoder.Decode(&dto, r.Form)
	if err != nil {
		return dto, getInvalidInputError(err)
	}
	return dto, nil
}

func getInvalidInputError(err error) *model.AppError {
	return &model.AppError{
		Type:    string(model.InvalidInputError),
		Context: err,
		Args: map[string]string{
			"message": err.Error(),
		},
	}
}
