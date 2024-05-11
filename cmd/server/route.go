package main

import (
	"github.com/go-chi/chi/v5"
	m "github.com/go-chi/chi/v5/middleware"
	"io"
	"log/slog"
	"net/http"
	"path/filepath"
	"time"
	"webapp/pkg/env"
	"webapp/pkg/middleware"
)

func index(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir(filepath.Join(env.AppWd, "ui")))
	slog.Info("PAGE", "method", r.Method, "path", r.URL.Path)
	fs.ServeHTTP(w, r)
}

func OK() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(56 * time.Second)
		_, _ = io.WriteString(w, "OK")
	})
}

func (app *application) routes() *chi.Mux {

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.ApplicationContextM)
	r.Use(m.RealIP)
	r.Use(middleware.Logger)

	r.Use(m.Recoverer)

	r.Use(m.Timeout(60 * time.Second))

	r.Route("/users", func(r chi.Router) {
		r.Get("/", app.allUsers)
	})

	return r
}
