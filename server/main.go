package main

import (
	"boilerplate/config"
	"boilerplate/env"
	"boilerplate/route"
	"boilerplate/template"
	"log/slog"
	"net/http"
	"os"
)

func init() {
	env.LoadEnvs()
	config.InitLogger()
	template.LoadTemplates()
	config.New()
}

func main() {
	template.Templates["version.tmpl"].Execute(os.Stdout, os.Getenv("APP_VERSION"))
	config.Logger.Info("Hello World!")
	db := config.MemDatabase(config.AppConfiguration)
	config.Logger.Info("created database connection", "db", db)

	mux := route.NewRoutes(db)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		slog.Error("Error starting server", "error", err)
		os.Exit(1)
	}
}
