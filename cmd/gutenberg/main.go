package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/byoc-io/gutenberg/authoring"
	"github.com/go-kit/kit/log"

	"github.com/byoc-io/gutenberg/pkg/storage/memory"
	"golang.org/x/net/context"
)

const (
	defaultPort = "8080"
)

func main() {
	var (
		addr     = envString("PORT", defaultPort)
		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")

		ctx = context.Background()
	)

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = &serializedLogger{Logger: logger}
	logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)

	// Setup repositories
	repository := memory.NewRepository()

	var as authoring.Service
	as = authoring.NewService(repository)
	as = authoring.NewLoggingService(log.NewContext(logger).With("component", "authoring"), as)

	httpLogger := log.NewContext(logger).With("component", "http")
	mux := http.NewServeMux()
	mux.Handle("/authoring/v1/", authoring.MakeHandler(ctx, as, httpLogger))

	http.Handle("/", accessControl(mux))

	errors := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errors <- http.ListenAndServe(*httpAddr, nil)
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errors <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errors)

}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

type serializedLogger struct {
	mtx sync.Mutex
	log.Logger
}

func (l *serializedLogger) Log(keyvals ...interface{}) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	return l.Logger.Log(keyvals...)
}
