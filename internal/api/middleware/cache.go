package middleware

import (
	"fmt"
	"net/http"
)

func NoCacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		next.ServeHTTP(writer, request)
	})
}

func GetCacheMiddleware(ttl int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", ttl))
			next.ServeHTTP(writer, request)
		})
	}
}
