package main

import (
	"log/slog"
	"os"

	"github.com/icomp-projects/tgconn/internal/api"
	"github.com/icomp-projects/tgconn/internal/env"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg := api.Config{
		Addr: env.GetString("ADDR", ":5555"),
	}

	app := api.New(cfg, logger)

	router := app.Mount()

	app.Logger.Info("starting server", "addr", cfg.Addr)

	err := app.Run(router)

	if err != nil {
		app.Logger.Error("failed to start server", "error", err)
	}

}
