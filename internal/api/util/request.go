package util

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func DeserializeAndValidateBody[T any](r *http.Request, validate *validator.Validate) (T, error) {
	dto, err := DeserializeBody[T](r)
	if err != nil {
		return dto, err
	}

	err = validate.Struct(dto)
	return dto, err
}

func DeserializeBody[T any](r *http.Request) (T, error) {
	var dto T
	err := json.NewDecoder(r.Body).Decode(&dto)
	return dto, err
}
