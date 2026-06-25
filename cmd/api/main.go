package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger slog.Logger
}

func main() {
	addr := flag.String("Addr", ":8080", "network port")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &application{logger: *logger}

	logger.Info("Starting server", "addr", addr)
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)

}
