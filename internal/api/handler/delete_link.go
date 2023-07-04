package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"gotiny/internal/api/util"
	"gotiny/internal/core/usecase"
)

// swagger:parameters deleteLink
type deleteLinkParams struct {
	// in: path
	LinkId string `json:"linkId"`
	// in: query
	Token string `json:"token"`
}

// swagger:route DELETE /api/v1/link/{linkId} link deleteLink
//
// # Delete link
//
// Delete link with given id.
//
// Responses:
//
//	204: emptyResponse
//	400: errorResponse
//	401: errorResponse
//	404: errorResponse
//	500: errorResponse
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
	err := h.deleteLink.Call(request.Context(), id, token)
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
