package main

import (
	"log/slog"

	"github.com/icomp-projects/tgconn/internal/api"
	"github.com/icomp-projects/tgconn/internal/env"
)

func main() {

	cfg := api.Config{
		Addr: env.GetString("ADDR", ":5555"),
	}

	app := api.New(cfg)

	router := app.Mount()

	slog.Info("starting server", "addr", cfg.Addr)

	err := app.Run(router)

	if err != nil {
		slog.Error("failed to start server", "error", err)
	}

}
