package middleware

import (
	"log"
	"net/http"
)

// CORSMiddleware adds CORS headers to the response.
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("CORS middleware executed")
		// Set the CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")                                // Allow any origin
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") // Allowed methods
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		// w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, origin") // Allowed headers

		// If it's a preflight OPTIONS request, stop here
		if r.Method == "OPTIONS" {
			// w.WriteHeader(http.StatusOK)
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
