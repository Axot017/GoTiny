package handler

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"

	"gotiny/internal/api/middleware"
	"gotiny/internal/api/util"
)

type AjaxHomePageHandler struct {
	template *template.Template
}

func NewAjaxHomePageHandler(
	template *template.Template,
) *AjaxHomePageHandler {
	return &AjaxHomePageHandler{
		template: template,
	}
}

func (h *AjaxHomePageHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	util.WriteTemplate(request, writer, h.template, "home_page.html", nil)
}

func (h *AjaxHomePageHandler) Register(router chi.Router) {
	router.With(middleware.GetCacheMiddleware(86400)).Get("/", h.ServeHTTP)
}
