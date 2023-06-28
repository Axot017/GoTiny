package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"gotiny/internal/api/util"
	"gotiny/internal/core/model"
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
	link, err := h.createShortLink.Call(request.Context(), id, token)
	if err != nil {
		util.WriteError(writer, err)
		return
	}

	if link == nil {
		util.WriteError(writer, model.NewNotFoundError())
		return
	}

	util.WriteResponseJson(writer, link)
}

func (h *GetLinkDetails) Path() string {
	return "/v1/link/{linkId:[a-zA-Z0-9]{1,}}"
}

func (h *GetLinkDetails) Method() string {
	return http.MethodGet
}
