package handler

import (
	"encoding/json"
	"net/http"

	"gotiny/internal/core/model"
	"gotiny/internal/core/usecase"
)

type CreateLinkHandler struct {
	createShortLink *usecase.CreateShortLink
}

func NewCreateLinkHandler(createShortLink *usecase.CreateShortLink) *CreateLinkHandler {
	return &CreateLinkHandler{createShortLink}
}

func (h *CreateLinkHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	config := model.LinkConfig{
		Protocol: "http",
		Host:     "localhost:8080",
	}
	link, err := h.createShortLink.Call("https://www.google.com", config)
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
