package middleware

import (
	"net/http"
)

// OnlyMethod ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func OnlyMethod(m ...string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			contains := false
			for _, m := range m {
				if m == r.Method {
					contains = true
					break
				}
			}

			if !contains {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			f(w, r)
		}
	}
}
