package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

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

func (h *HealthHandler) Register(router chi.Router) {
	router.Get("/api/health", h.ServeHTTP)
}
