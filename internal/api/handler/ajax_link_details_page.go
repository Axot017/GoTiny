package handler

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"

	"gotiny/internal/api/middleware"
	"gotiny/internal/api/util"
	"gotiny/internal/core/model"
)

type AjaxLinkDetailsPageHandler struct {
	template           *template.Template
	linkTokenValidator *middleware.LinkTokenValidator
}

func NewAjaxLinkDetailsPageHandler(
	template *template.Template,
	linkTokenValidator *middleware.LinkTokenValidator,
) *AjaxLinkDetailsPageHandler {
	return &AjaxLinkDetailsPageHandler{
		template:           template,
		linkTokenValidator: linkTokenValidator,
	}
}

func (h *AjaxLinkDetailsPageHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	link := request.Context().Value("link").(*model.Link)
	util.WriteTemplate(request, writer, h.template, "link_details_page.html", link)
}

func (h *AjaxLinkDetailsPageHandler) Register(router chi.Router) {
	router.
		With(h.linkTokenValidator.Handle, middleware.GetCacheMiddleware(86400)).
		Get("/link/{linkId}", h.ServeHTTP)
}
