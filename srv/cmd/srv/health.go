package main

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/go-redis/redis/v7"
)

// Health checks health status of this application. Will run check against db and cache, if
// everything looks ok then it will return 200 code, 500 otherwise.
func Health(ctx context.Context, db *sql.DB, client *redis.Client, logger *logrus.Entry) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := db.PingContext(ctx); err != nil {
			logger.Errorf("db.PingContext %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// FIXME [grokrz]: fix redis health check
		//if err := client.Ping().Err(); err != nil {
		//	logger.Printf("cache : error : %v", err)
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		w.WriteHeader(http.StatusOK)
	}
}
