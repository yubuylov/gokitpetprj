package main

import (
	"net/http"

	app_config "github.com/yubuylov/gokitpetprj/storage/config"
	app_logs "github.com/yubuylov/gokitpetprj/storage/loggers"
	app_mertics "github.com/yubuylov/gokitpetprj/storage/metrics"

	"github.com/yubuylov/gokitpetprj/storage/api"
	"github.com/yubuylov/gokitpetprj/storage/backend"
	"os/signal"
	"fmt"
	"os"
	"syscall"
	"golang.org/x/net/context"
)

var App = struct {
	Cfg     app_config.AppConfig
	Logs    app_logs.AppLogs
	Metrics app_mertics.AppMetrics
}{}

func main() {
	App.Cfg = app_config.Load()
	App.Logs = app_logs.Load(App.Cfg)
	App.Metrics = app_mertics.Load(App.Cfg)

	ctx := context.Background()
	mux := http.NewServeMux()

	storage := backend.NewStorage(App.Cfg, &App.Logs)
	bs := backend.InitService(storage, App.Logs, App.Metrics)

	mux.Handle("/api/v1/", api.Handler(ctx, bs, App.Logs.Access))

	errCh := make(chan error, 2)
	runServer(mux, errCh)
	waitSyscall(errCh)
	App.Logs.Error.Log("terminated", <-errCh)
}

func runServer(h http.Handler, errCh chan error) {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	}))

	go func() {
		App.Logs.Info.Log("listen", "HTTP", "addr", App.Cfg.Server.Listen)
		errCh <- http.ListenAndServe(App.Cfg.Server.Listen, nil)
	}()
}

func waitSyscall(errCh chan error) {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errCh <- fmt.Errorf("%s", <-c)
	}()
}