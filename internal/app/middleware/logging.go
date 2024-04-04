package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logs the HTTP request details and the time taken to process the request
// It has been enchanced to log additional CORS-related information
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // Start the timer

		// Check if the request is a CORS preflight request
		isPreflight := r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != ""
		origin := r.Header.Get("Origin")

		// Log the incoming request with CORS details
		if isPreflight {
			log.Printf("CORS Preflight from origin %s for %s %s", origin, r.Method, r.RequestURI)
		} else if origin != "" {
			log.Printf("CORS request from origin %s: Started %s %s", origin, r.Method, r.RequestURI)
		} else {
			log.Printf("Started %s %s", r.Method, r.RequestURI)
		}

		next.ServeHTTP(w, r) // Pass control to the next handler

		// Log the time take to process the request, including CORS preflight indication if applicable
		if isPreflight {
			log.Printf("CORS Preflight from origin %s for %s %s completed in %v", origin, r.Method, r.RequestURI, time.Since(start))
		} else {
			log.Printf("Completed %s in %v", r.RequestURI, time.Since(start))
		}
	})
}
