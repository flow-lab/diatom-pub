package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/go-redis/redis/v7"
)

// Health checks health status of this application. Will run check against db and cache, if
// everything looks ok then it will return 200 code, 500 otherwise.
func Health(ctx context.Context, db *sql.DB, client *redis.Client, log *log.Logger) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.PingContext(ctx); err != nil {
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
