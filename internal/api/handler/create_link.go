package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"

	"gotiny/internal/api/dto"
	"gotiny/internal/api/util"
	"gotiny/internal/core/model"
	"gotiny/internal/core/usecase"
)

type CreateLinkHandler struct {
	createShortLink *usecase.CreateShortLink
	validate        *validator.Validate
}

func NewCreateLinkHandler(
	createShortLink *usecase.CreateShortLink,
	validate *validator.Validate,
) *CreateLinkHandler {
	return &CreateLinkHandler{createShortLink, validate}
}

func (h *CreateLinkHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	create_link_dto, err := util.DeserializeAndValidateBody[dto.CreateLinkDto](request, h.validate)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	config := model.LinkConfig{
		MaxHits:    create_link_dto.MaxHits,
		ValidUntil: create_link_dto.ValidUntil,
		Host:       request.Host,
	}
	link, err := h.createShortLink.Call(create_link_dto.Link, config)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(link)
}

func (h *CreateLinkHandler) Path() string {
	return "/v1/link"
}

func (h *CreateLinkHandler) Method() string {
	return http.MethodPost
}
