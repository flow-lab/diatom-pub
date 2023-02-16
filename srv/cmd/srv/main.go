package main

import (
	"context"
	"database/sql"
	"expvar"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/flow-lab/diatom-pub/internal/db"
	"github.com/flow-lab/diatom-pub/internal/handler"
	"github.com/flow-lab/diatom-pub/internal/middleware"
	"github.com/go-redis/redis/v7"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	logger := log.New(os.Stdout, fmt.Sprintf("diatom-pub : (%s, %s) : ", version, short(commit)), log.LstdFlags|log.Lmicroseconds|log.Lshortfile|log.Ldate)
	if err := run(logger); err != nil {
		logger.Printf("error : %s", err)
		os.Exit(1)
	}
}

func run(logger *log.Logger) error {
	expvar.NewString("version").Set(version)
	expvar.NewString("commit").Set(commit)
	expvar.NewString("date").Set(date)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		logger.Fatal(errors.New("missing required parameter PORT"))
	}

	templateDir := os.Getenv("TEMPLATE_DIR")
	if templateDir == "" {
		templateDir = "template/"
	}

	staticDir := os.Getenv("STATIC_DIR")
	if staticDir == "" {
		staticDir = "static/"
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	logger.Printf("gomaxprocs : %d", runtime.GOMAXPROCS(0))

	logger.Println("api server : initializing ")
	readTimeout := 60 * time.Second
	writeTimeout := 60 * time.Second
	apiSrv := http.Server{
		Addr:         ":" + port,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	// connect to db
	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		logger.Fatal(errors.New("missing required parameter DB_HOST"))
	}
	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		logger.Fatal(errors.New("missing required parameter DB_PORT"))
	}
	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		logger.Fatal(errors.New("missing required parameter DB_NAME"))
	}
	dbUsername, ok := os.LookupEnv("DB_USERNAME")
	if !ok {
		logger.Fatal(errors.New("missing required parameter DB_USERNAME"))
	}
	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		logger.Fatal(errors.New("missing required parameter DB_PASSWORD"))
	}

	dbSSLMode := os.Getenv("DB_SSLMODE")

	// connect to db
	dbClient, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", dbHost, dbPort, dbName, dbUsername, dbPassword, dbSSLMode))
	if err != nil {
		logger.Fatal(errors.Wrap(err, "sql.Open"))
	}

	// print db stats
	logger.Printf("db stats : %+v", dbClient.Stats())

	// queries
	queries := db.New(dbClient)

	// connect to redis
	redisHost, ok := os.LookupEnv("REDIS_HOST")
	if !ok {
		logger.Fatal(errors.New("missing required parameter REDIS_HOST"))
	}
	redisPort, ok := os.LookupEnv("REDIS_PORT")
	if !ok {
		logger.Fatal(errors.New("missing required parameter REDIS_PORT"))
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})

	// print redis stats
	logger.Printf("redis stats : %+v", redisClient.PoolStats())

	ctx := context.Background()

	// start api server
	go func(ctx context.Context) {
		logger.Printf("api server : listening on %s", apiSrv.Addr)
		http.HandleFunc("/health", middleware.Chain(
			Health(ctx, dbClient, redisClient, logger),
			middleware.OnlyMethod("GET")),
		)
		http.HandleFunc("/api.yaml", middleware.Chain(
			APIDefinition(templateDir, logger),
			middleware.OnlyMethod("GET"),
			middleware.Sec(),
			middleware.CORS(),
			middleware.Logging(logger)),
		)
		http.HandleFunc("/authors/", middleware.Chain(
			handler.GetAuthor(logger, queries),
			middleware.OnlyMethod("GET"),
			middleware.Sec(),
			middleware.CORS(),
			middleware.Logging(logger)),
		)
		serverErrors <- apiSrv.ListenAndServe()
	}(ctx)

	// listen to all signal from os
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		return err
	case sig := <-shutdown:
		timeout := 5 * time.Second
		logger.Printf("got %v. Start graceful shutdown with timeout %s", sig, timeout)

		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		// wait for timeout
		time.Sleep(timeout)

		// Asking listener to shut down and load shed.
		err := apiSrv.Shutdown(ctx)
		if err != nil {
			logger.Printf("graceful shutdown timeout in %v : %v", timeout, err)
			err = apiSrv.Close()
		}

		switch {
		case sig == syscall.SIGSTOP:
			logger.Printf("integrity error : %v", sig)
			return errors.New("integrity issue caused shutdown")
		case err != nil:
			logger.Printf("error : %v", err)
			return err
		}
	}

	return nil
}

func short(s string) string {
	if len(s) > 7 {
		return s[0:7]
	}
	return s
}
