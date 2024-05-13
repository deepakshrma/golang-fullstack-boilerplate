package main

import (
	"github.com/go-chi/chi/v5"
	m "github.com/go-chi/chi/v5/middleware"
	"time"
	"webapp/pkg/middleware"
)

func (app *application) routes() *chi.Mux {

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.ApplicationContextM)
	r.Use(m.RealIP)
	r.Use(middleware.Logger)

	r.Use(m.Recoverer)

	r.Use(m.Timeout(60 * time.Second))
	r.Get("/", app.Home)

	sh := &staticHandler{"ui", ""}
	r.Get("/static/*", sh.ServeHTTP)

	r.Route("/api/users", func(r chi.Router) {
		r.Get("/", app.allUsers)
	})
	r.Route("/api/todos", func(r chi.Router) {
		r.Get("/", app.allTodos)
	})

	return r
}
