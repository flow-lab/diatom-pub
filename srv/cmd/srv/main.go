package main

import (
	"context"
	"expvar"
	"github.com/flow-lab/diatom-pub/internal/cache"
	"github.com/flow-lab/dlog"
	"github.com/sirupsen/logrus"
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
	utils "github.com/flow-lab/utils"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	logger := dlog.NewLogger(&dlog.Config{
		AppName:      "diatom-pub",
		Level:        utils.EnvOrDefault("LOG_LEVEL", "debug").(string),
		Version:      version,
		Commit:       utils.Short(commit),
		Build:        date,
		ReportCaller: true,
	})
	if err := run(logger); err != nil {
		logger.Errorf("run %s", err)
		os.Exit(1)
	}
}

func run(logger *logrus.Entry) error {
	expvar.NewString("version").Set(version)
	expvar.NewString("commit").Set(commit)
	expvar.NewString("date").Set(date)

	// catch the panic
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("panic : %s", r)
		}
	}()

	var (
		port        = utils.MustGetEnv("PORT")
		templateDir = utils.EnvOrDefault("TEMPLATE_DIR", "template/").(string)
	)

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	logger.Infof("gomaxprocs %d", runtime.GOMAXPROCS(0))
	logger.Infof("api server initializing...")
	readTimeout := 30 * time.Second
	writeTimeout := 30 * time.Second
	apiSrv := http.Server{
		Addr:         ":" + port,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	// connect to db
	dbClient, err := db.ConnectTCPSocket()
	if err != nil {
		logger.Fatal(errors.Wrap(err, "sql.Open"))
	}

	// print db stats
	logger.Infof("db stats : %+v", dbClient.Stats())

	// queries
	queries := db.New(dbClient)

	redisClient, err := cache.NewClient()
	if err != nil {
		logger.Fatal(errors.Wrap(err, "cache.NewClient"))
	}

	// print redis stats
	logger.Infof("redis stats : %+v", redisClient.PoolStats())

	ctx := context.Background()

	// start api server
	go func(ctx context.Context) {
		logger.Infof("api server : listening on %s", apiSrv.Addr)
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
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSTOP)

	select {
	case err := <-serverErrors:
		return err
	case sig := <-shutdown:
		timeout := 10 * time.Second
		logger.Infof("got %v. Start graceful shutdown with timeout %s", sig, timeout)
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		// Asking listener to shut down and load shed.
		err := apiSrv.Shutdown(ctx)
		if err != nil {
			logger.Infof("graceful shutdown timeout in %v : %v", timeout, err)
			err = apiSrv.Close()
		}

		switch {
		case sig == syscall.SIGSTOP:
			logger.Errorf("integrity error : %v", sig)
			return errors.New("integrity issue caused shutdown")
		case err != nil:
			logger.Errorf("error %v", err)
			return err
		}
	}

	return nil
}
