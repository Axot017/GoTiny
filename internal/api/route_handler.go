package api

import "net/http"

type RouteHandler interface {
	http.Handler

	Path() string

	Method() string
}
