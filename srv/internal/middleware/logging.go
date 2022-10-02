package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging logs all requests with its path and the processing time
func Logging(log *log.Logger) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rr := resWriterWrapper(w)
			defer func(rr *res) {
				log.Println(r.URL.Path, rr.Status(), time.Since(start))
			}(rr)
			f(rr, r)
		}
	}
}

type res struct {
	http.ResponseWriter
	code int
}

// Wraps ResponseWriter to get response code.
func resWriterWrapper(w http.ResponseWriter) *res {
	return &res{w, http.StatusOK}
}

func (r *res) Status() int {
	return r.code
}

func (r *res) Header() http.Header {
	return r.ResponseWriter.Header()
}

func (r *res) Write(data []byte) (int, error) {
	return r.ResponseWriter.Write(data)
}

func (r *res) WriteHeader(statusCode int) {
	r.code = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
