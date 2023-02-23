package main

import (
	"github.com/flow-lab/dlog"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAPIDefinition(t *testing.T) {
	t.Run("should server API def", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		var h http.HandlerFunc
		logger := dlog.NewLogger(&dlog.Config{
			AppName: "test",
		})
		h = APIDefinition("../../template/", logger)
		h.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		res := rec.Body.String()

		if !strings.Contains(res, "openapi: 3.0") {
			t.Errorf("%s does not return expected file", res)
		}
	})
}
