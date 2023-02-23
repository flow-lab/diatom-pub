package main

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"strings"
)

// APIDefinition serves oas3 api definition.
func APIDefinition(path string, logger *logrus.Entry) func(w http.ResponseWriter, _ *http.Request) {
	t, err := template.ParseFiles(path + "api.yaml")
	if err != nil {
		logger.Fatalf("unable to load template api.yaml, err: %v", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		url := ""
		if strings.Contains(r.Host, "localhost") {
			url = fmt.Sprintf("http://%s", r.Host)
		} else {
			url = fmt.Sprintf("https://%s", r.Host)
		}
		var data = struct {
			SrvUrl string
		}{
			SrvUrl: url,
		}

		var buf bytes.Buffer
		if err := t.Execute(&buf, data); err != nil {
			logger.Errorf("t.Execute %s", errors.Wrap(err, "api template exec"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/vnd.yaml")
		_, err := w.Write(buf.Bytes())
		if err != nil {
			logger.Errorf("w.Write %v", err)
		}
	}
}
