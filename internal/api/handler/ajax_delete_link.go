package handler

import (
	"net/http"

	"gotiny/internal/api/middleware"
	"gotiny/internal/core/model"
	"gotiny/internal/core/usecase"
)

type AjaxDeleteLinkHandler struct {
	deleteLink         *usecase.DeleteLink
	linkTokenValidator *middleware.LinkTokenValidator
}

func NewAjaxDeleteLinkHandler(
	deleteLink *usecase.DeleteLink,
	linkTokenValidator *middleware.LinkTokenValidator,
) *AjaxDeleteLinkHandler {
	return &AjaxDeleteLinkHandler{
		deleteLink:         deleteLink,
		linkTokenValidator: linkTokenValidator,
	}
}

func (h *AjaxDeleteLinkHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	link := request.Context().Value("link").(*model.Link)
	_ = h.deleteLink.Call(request.Context(), link.Id)

	writer.Header().Set("HX-Redirect", "/")
}

func (h *AjaxDeleteLinkHandler) Path() string {
	return "/link/{linkId:[a-zA-Z0-9]{1,}}"
}

func (h *AjaxDeleteLinkHandler) Method() string {
	return http.MethodDelete
}

func (h *AjaxDeleteLinkHandler) Middlewares() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		h.linkTokenValidator.Handle,
	}
}