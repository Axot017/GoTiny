package api

import (
	"github.com/go-chi/chi/v5"
)

type RouteHandler interface {
	Register(router chi.Router)
}
