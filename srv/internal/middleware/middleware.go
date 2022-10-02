package middleware

import (
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(f http.HandlerFunc, ms ...Middleware) http.HandlerFunc {
	for _, m := range ms {
		f = m(f)
	}
	return f
}
