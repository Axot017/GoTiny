package middleware

import "net/http"

type AppMiddleware interface {
	Handle(next http.Handler) http.Handler
}

type WithMiddlewares interface {
	Middlewares() []AppMiddleware
}

func GetMiddlewaresToAttach(handler interface{}) []AppMiddleware {
	if h, ok := handler.(WithMiddlewares); ok {
		return h.Middlewares()
	}
	return []AppMiddleware{}
}
