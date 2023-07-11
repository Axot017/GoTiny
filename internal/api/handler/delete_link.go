package handler

import (
	"net/http"

	"gotiny/internal/api/middleware"
	"gotiny/internal/api/util"
	"gotiny/internal/core/model"
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
	deleteLink         *usecase.DeleteLink
	linkTokenValidator *middleware.LinkTokenValidator
}

func NewDeleteLinkHandler(
	deleteLink *usecase.DeleteLink,
	linkTokenValidator *middleware.LinkTokenValidator,
) *DeleteLinkHandler {
	return &DeleteLinkHandler{
		deleteLink:         deleteLink,
		linkTokenValidator: linkTokenValidator,
	}
}

func (h *DeleteLinkHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	link := request.Context().Value("link").(*model.Link)
	err := h.deleteLink.Call(request.Context(), link.Id)
	if err != nil {
		util.WriteError(writer, err)
		return
	}

	util.WriteResponseJson(writer, nil, http.StatusNoContent)
}

func (h *DeleteLinkHandler) Path() string {
	return "/api/v1/link/{linkId:[a-zA-Z0-9]{1,}}"
}

func (h *DeleteLinkHandler) Method() string {
	return http.MethodDelete
}

func (h *DeleteLinkHandler) Middlewares() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		h.linkTokenValidator.Handle,
	}
}
