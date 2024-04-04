package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
)

// ErrorHandlerFunc is a function type that matches our handlers' signature.
type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request) error

// ServeHTTP makes ErrorHandlerFunc fit http.Handler interface.
// This method calls the ErrorHandlerFunc and handles errors.
func (fn ErrorHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		// Here we can log the error, send a specific status code, etc.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ErrorHandler catches and handles errors and panics in HTTP handlers.
func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v\n%s", err, debug.Stack())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
