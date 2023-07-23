package middleware

import (
	"context"
	"net/http"

	"github.com/segmentio/ksuid"
)

type IdCookieMiddleware struct{}

func NewIdCookieMiddleware() *IdCookieMiddleware {
	return &IdCookieMiddleware{}
}

func (m *IdCookieMiddleware) Handle(next http.Handler) http.Handler {
	fn := func(writer http.ResponseWriter, request *http.Request) {
		id_cookie, err := request.Cookie("user_id")
		if err != nil {
			id_cookie = &http.Cookie{
				Name:  "user_id",
				Value: ksuid.New().String(),
			}
			http.SetCookie(writer, id_cookie)
		}

		ctx := context.WithValue(request.Context(), "user_id", &id_cookie.Value)

		next.ServeHTTP(writer, request.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
