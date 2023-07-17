package handler

import (
	"html/template"
	"net/http"

	"gotiny/internal/api/middleware"
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
	h.template.ExecuteTemplate(writer, "link_details_page.html", link)
}

func (h *AjaxLinkDetailsPageHandler) Path() string {
	return "/link/{linkId}"
}

func (h *AjaxLinkDetailsPageHandler) Method() string {
	return http.MethodGet
}

func (h *AjaxLinkDetailsPageHandler) Middlewares() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		h.linkTokenValidator.Handle,
	}
}
