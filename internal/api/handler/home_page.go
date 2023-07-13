package handler

import (
	"html/template"
	"net/http"
)

type HomePageHandler struct {
	template *template.Template
}

func NewHomePageHandler(
	template *template.Template,
) *HomePageHandler {
	return &HomePageHandler{
		template: template,
	}
}

func (h *HomePageHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	h.template.ExecuteTemplate(writer, "home.html", nil)
}

func (h *HomePageHandler) Path() string {
	return "/"
}

func (h *HomePageHandler) Method() string {
	return http.MethodGet
}
