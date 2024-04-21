package route

import (
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

func NewRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	commonMiddlewares := Combine(middleware.RequestContextID, middleware.Logger).Then(OK())

	mux.Handle("/todos", commonMiddlewares)

	// Index page http server
	mux.HandleFunc("/", index)

	return mux
}
