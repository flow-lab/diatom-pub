package handler

import (
	"database/sql"
	api "github.com/flow-lab/diatom-pub/internal/apimodel"
	"github.com/flow-lab/diatom-pub/internal/db"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

// GetAuthor serves GET /authors/:id
func GetAuthor(logger *logrus.Entry, queries *db.Queries) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/authors/"):]
		logger.Infof("id %s", id)

		authorID, err := uuid.Parse(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		authorDB, err := queries.GetAuthor(r.Context(), authorID)
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			logger.Errorf("GetAuthor error %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		authorAPI := api.NewAuthor()
		idStr := authorDB.ID.String()
		authorAPI.Id = &idStr
		authorAPI.Name = &authorDB.Name

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json, err := authorAPI.MarshalJSON()
		if err != nil {
			logger.Errorf("authorAPI.MarshalJSON %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(json)
		if err != nil {
			logger.Printf("w.Write %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
