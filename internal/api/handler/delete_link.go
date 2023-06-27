package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"gotiny/internal/api/util"
	"gotiny/internal/core/usecase"
)

type DeleteLinkHandler struct {
	deleteLink *usecase.DeleteLink
}

func NewDeleteLinkHandler(deleteLink *usecase.DeleteLink) *DeleteLinkHandler {
	return &DeleteLinkHandler{
		deleteLink: deleteLink,
	}
}

func (h *DeleteLinkHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "linkId")
	token := request.URL.Query().Get("token")
	err := h.deleteLink.Call(id, token)
	if err != nil {
		util.WriteError(writer, err)
		return
	}

	util.WriteResponseJson(writer, nil, http.StatusNoContent)
}

func (h *DeleteLinkHandler) Path() string {
	return "/v1/link/{linkId:[a-zA-Z0-9]{1,}}"
}

func (h *DeleteLinkHandler) Method() string {
	return http.MethodDelete
}
