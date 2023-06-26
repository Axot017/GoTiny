package handler

import "net/http"

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}

func (h *HealthHandler) Path() string {
	return "/health"
}

func (h *HealthHandler) Method() string {
	return http.MethodGet
}
