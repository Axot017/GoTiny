package handler

import (
	"net/http"

	"gotiny/internal/core/usecase"
)

type CreateLinkHandler struct {
	createShortLink *usecase.CreateShortLink
}

func NewCreateLinkHandler(createShortLink *usecase.CreateShortLink) *CreateLinkHandler {
	return &CreateLinkHandler{createShortLink}
}

func (h *CreateLinkHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *CreateLinkHandler) Path() string {
	return "/link"
}

func (h *CreateLinkHandler) Method() string {
	return http.MethodPost
}
