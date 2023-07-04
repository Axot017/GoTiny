package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"gotiny/internal/api/util"
	"gotiny/internal/core/model"
	"gotiny/internal/core/usecase"
)

// swagger:route GET /{linkId} redirect
//
// # Redirect
//
// This will redirect to the original URL.
//
// Responses:
//
//	302: emptyResponse
//	500: errorResponse
type RedirectHandler struct {
	hitLink *usecase.HitLink
}

func NewRedirectHandler(hitLink *usecase.HitLink) *RedirectHandler {
	return &RedirectHandler{
		hitLink: hitLink,
	}
}

func (h *RedirectHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "linkId")
	url, err := h.hitLink.Call(request.Context(), id)
	if err != nil {
		util.WriteError(writer, err)
		return
	}
	if url == nil {
		util.WriteError(writer, model.NewNotFoundError())
		return
	}

	http.Redirect(writer, request, *url, http.StatusMovedPermanently)
}

func (h *RedirectHandler) Path() string {
	return "/{linkId:[a-zA-Z0-9]{1,}}"
}

func (h *RedirectHandler) Method() string {
	return http.MethodGet
}
