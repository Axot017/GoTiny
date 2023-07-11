package middleware

import "net/http"

type AppMiddleware interface {
	Handle(next http.Handler) http.Handler
}

type WithMiddlewares interface {
	Middlewares() []func(http.Handler) http.Handler
}

func GetMiddlewaresToAttach(handler interface{}) []func(http.Handler) http.Handler {
	if h, ok := handler.(WithMiddlewares); ok {
		return h.Middlewares()
	}
	return []func(http.Handler) http.Handler{}
}
