package api

import (
	"log"
	"net/http"
	"time"
)

type MiddlewareFunc func(http.Handler) http.Handler


func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		log.Printf("%s | %s - %s", currentTime.Format(time.RFC3339), r.RequestURI, r.Method)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
