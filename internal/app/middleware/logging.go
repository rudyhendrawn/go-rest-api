package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logs the HTTP request details and the time taken to process the request
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the incoming request
		log.Printf("Started %s %s", r.Method, r.RequestURI)

		next.ServeHTTP(w, r) // Pass control to the next handler

		// Log the time take to process the request
		log.Printf("Completed %s in %v", r.RequestURI, time.Since(start))
	})
}
