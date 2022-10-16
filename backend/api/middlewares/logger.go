package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s -> %s | %s | Body: %dB\n", time.Now(), r.Method, r.RequestURI, r.Host, r.Proto, r.ContentLength)
		next.ServeHTTP(w, r)
	})
}
