package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"gotiny/internal/api/middleware"
	"gotiny/internal/api/util"
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
	err := h.deleteLink.Call(request.Context(), link.Id)
	if err != nil {
		util.WriteAjaxError(writer, err)
		return
	}

	writer.Header().Set("HX-Redirect", "/")
}

func (h *AjaxDeleteLinkHandler) Register(router chi.Router) {
	router.With(h.linkTokenValidator.Handle).Delete("/ajax/link/{linkId}", h.ServeHTTP)
}
