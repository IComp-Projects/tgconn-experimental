package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type application struct {
	config Config
}

type Config struct {
	Addr string
}

func New(cfg Config) *application {
	return &application{
		config: cfg,
	}
}

func (app *application) Mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/webhook", app.HandleUpdates)
	})

	return r
}

func (app *application) Run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.Addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	return srv.ListenAndServe()
}
