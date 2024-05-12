package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"webapp/pkg/config"
	"webapp/pkg/env"
	"webapp/pkg/repository"
	"webapp/pkg/repository/repo"
	"webapp/templates"
)

type application struct {
	//db  *sql.DB
	DSN    string
	db     repository.DatabaseRepo
	L      *slog.Logger
	config *config.AppConfig
}

var app application

func init() {
	env.LoadEnvs()
	app.L = config.NewLogger()
	templates.LoadTemplates()
	cnf, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}
	app.config = cnf
}

func main() {
	app.DSN = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	err := templates.Templates["version.tmpl"].Execute(os.Stdout, os.Getenv("APP_VERSION"))
	if err != nil {
		return
	}

	app.L.Info("Hello World!")
	db, err := app.connectToDB()
	if err != nil {
		slog.Error("Error connecting to DB", err)
	}
	app.L.Info("created database connection")

	app.db = &repo.PostgresDBRepo{DB: db}
	mux := app.routes()
	app.L.Info("server is running", "port", app.config.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", app.config.Port), mux)
	if err != nil {
		log.Fatal("Error starting server", "error", err)
	}
}
