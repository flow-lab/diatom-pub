package middleware

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORS(t *testing.T) {
	t.Run("should return OK for OPTIONS", func(t *testing.T) {
		req, err := http.NewRequest("OPTIONS", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		h := Chain(func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte("ok"))
		}, CORS())

		h.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("response code was %v instead of %d", rec.Code, http.StatusUnauthorized)
		}
	})

	t.Run("should add CORS headers", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		h := Chain(func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte("ok"))
		}, CORS())

		h.ServeHTTP(rec, req)

		assert.Equal(t, "*", rec.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, "OPTIONS,POST", rec.Header().Get("Access-Control-Allow-Methods"))
		assert.Equal(t, "content-type,api-key", rec.Header().Get("Access-Control-Allow-Headers"))
	})
}
