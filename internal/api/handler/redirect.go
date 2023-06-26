package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"gotiny/internal/core/usecase"
)

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
	url, err := h.hitLink.Call(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if url == nil {
		writer.WriteHeader(http.StatusNotFound)
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
