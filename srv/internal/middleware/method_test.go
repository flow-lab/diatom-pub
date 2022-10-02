package middleware

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ok(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("ok"))
}

func TestOnlyMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	check(t, err)

	rec := httptest.NewRecorder()
	h := Chain(ok, OnlyMethod("GET"))

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestOnlyMethod2(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	check(t, err)

	rec := httptest.NewRecorder()
	h := Chain(ok, OnlyMethod("GET"))

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
