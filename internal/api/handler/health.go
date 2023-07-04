package handler

import "net/http"

// swagger:route GET /api/health health
//
// # Health check
//
// This will check if the service is up and running.
//
// Responses:
//
//	200: emptyResponse
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
