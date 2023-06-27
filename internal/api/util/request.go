package util

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"

	"gotiny/internal/core/model"
)

func DeserializeAndValidateBody[T any](r *http.Request, validate *validator.Validate) (T, error) {
	dto, err := DeserializeBody[T](r)
	if err != nil {
		return dto, err
	}

	err = validate.Struct(dto)

	if err != nil {
		return dto, &model.AppError{
			Type:    string(model.InvalidInputError),
			Context: err,
			Args: map[string]string{
				"message": err.Error(),
			},
		}
	}
	return dto, nil
}

func DeserializeBody[T any](r *http.Request) (T, error) {
	var dto T
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return dto, &model.AppError{
			Type:    string(model.InvalidInputError),
			Context: err,
			Args: map[string]string{
				"message": err.Error(),
			},
		}
	}
	return dto, nil
}
