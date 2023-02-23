package middleware

import (
	"github.com/flow-lab/dlog"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogging(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	l := dlog.NewLogger(&dlog.Config{
		AppName: "test",
		Version: "test",
	})
	h := Chain(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}, Logging(l))

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
