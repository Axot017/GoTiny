package handler

import "net/http"

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *HealthHandler) Path() string {
	return "/health"
}

func (h *HealthHandler) Method() string {
	return http.MethodGet
}
