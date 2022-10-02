package main

import (
	"database/sql"
	"github.com/go-redis/redis/v7"
	"log"
	"net/http"
)

// Health checks health status of this application. Will run check against db and cache, if
// everything looks ok then it will return 200 code, 500 otherwise.
func Health(log *log.Logger, db *sql.DB, client *redis.Client) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			log.Printf("db : error : %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := client.Ping().Err(); err != nil {
			log.Printf("cache : error : %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
