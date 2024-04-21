package route

import (
	"boilerplate/config"
	"boilerplate/controller"
	"boilerplate/env"
	"boilerplate/route/middleware"
	"log/slog"
	"net/http"
	"path/filepath"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir(filepath.Join(env.AppWd, "ui")))
	slog.Info("PAGE", "method", r.Method, "path", r.URL.Path)
	fs.ServeHTTP(w, r)
}
func OK() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(56 * time.Millisecond)
		w.Write([]byte("OK"))
	})
}

func NewRoutes(db *config.Database) *http.ServeMux {
	mux := http.NewServeMux()

	uc := controller.NewUsersHandler(db)
	commonMiddlewares := Combine(middleware.RequestContextID, middleware.Logger, middleware.Recovery).Then(OK())
	usersMiddlewares := Combine(middleware.RequestContextID, middleware.Logger, middleware.Recovery).Then(http.HandlerFunc(uc.Users))

	mux.Handle("/ok", commonMiddlewares)
	mux.Handle("/users/", usersMiddlewares)

	// Index page http server
	mux.HandleFunc("/", index)

	return mux
}
