package handler

import (
	"html/template"
	"net/http"
)

type HomePageHandler struct{}

func NewHomePageHandler() *HomePageHandler {
	return &HomePageHandler{}
}

func (h *HomePageHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseGlob("web/templates/*")
	if err != nil {
		panic(err)
	}
	writer.WriteHeader(http.StatusOK)
	temp.ExecuteTemplate(writer, "home.html", nil)
}

func (h *HomePageHandler) Path() string {
	return "/"
}

func (h *HomePageHandler) Method() string {
	return http.MethodGet
}
