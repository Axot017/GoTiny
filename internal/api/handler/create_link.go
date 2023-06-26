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

func (h *CreateLinkHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dto, err := util.DeserializeAndValidateBody[dto.CreateLinkDto](r, h.validate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	config := model.LinkConfig{
		MaxHits:  dto.MaxHits,
		TtlInSec: dto.TtlInSec,
		Host:     r.Host,
	}
	link, err := h.createShortLink.Call(dto.Link, config)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(link)
}

func (h *CreateLinkHandler) Path() string {
	return "/link"
}

func (h *CreateLinkHandler) Method() string {
	return http.MethodPost
}
