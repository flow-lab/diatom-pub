package middleware

import (
	"net/http"
)

// Sec adds security headers.
func Sec() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// https://owasp.org/www-community/Security_Headers
			w.Header().Set("X-Content-Type-Options", "nosniff")
			f(w, r)
		}
	}
}
