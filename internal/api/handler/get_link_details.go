package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"gotiny/internal/core/usecase"
)

type GetLinkDetails struct {
	createShortLink *usecase.GetLinkDetails
}

func NewGetLinkDetails(createShortLink *usecase.GetLinkDetails) *GetLinkDetails {
	return &GetLinkDetails{
		createShortLink: createShortLink,
	}
}

func (h *GetLinkDetails) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "linkId")
	token := request.URL.Query().Get("token")
	link, err := h.createShortLink.Call(id, token)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if link == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(link)
}

func (h *GetLinkDetails) Path() string {
	return "/v1/link/{linkId:[a-zA-Z0-9]{1,}}"
}

func (h *GetLinkDetails) Method() string {
	return http.MethodGet
}
