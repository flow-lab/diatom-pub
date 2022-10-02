package middleware

import (
	"net/http"
)

// CORS enables Cross Origin Resource Sharing. For preflight requests it will return Ok and terminate middleware
// processing.
func CORS() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,POST")
			w.Header().Set("Access-Control-Allow-Headers", "content-type,api-key")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			f(w, r)
		}
	}
}
