package handler

import (
	"html/template"
	"net/http"

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

func (h *AjaxHomePageHandler) Path() string {
	return "/"
}

func (h *AjaxHomePageHandler) Method() string {
	return http.MethodGet
}
