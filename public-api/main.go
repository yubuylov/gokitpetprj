package main

import (
	"net/http"

	app_config "github.com/yubuylov/gokitpetprj/public-api/config"
	app_logs "github.com/yubuylov/gokitpetprj/public-api/loggers"
	app_mertics "github.com/yubuylov/gokitpetprj/public-api/metrics"

	"github.com/yubuylov/gokitpetprj/public-api/api"
	"github.com/yubuylov/gokitpetprj/public-api/backend"
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

	bs := backend.InitService(App.Cfg, App.Logs, App.Metrics)

	mux.Handle("/entities", api.Handler(ctx, bs, App.Logs.Access, App.Cfg))

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