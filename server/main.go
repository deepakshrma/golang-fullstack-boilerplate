package main

import (
	"boilerplate/config"
	"boilerplate/env"
	"boilerplate/template"
	"fmt"
	"os"
)

func init() {
	env.LoadEnvs()
	template.LoadTemplates()
	config.New()
}

func main() {
	template.Templates["version.tmpl"].Execute(os.Stdout, os.Getenv("APP_VERSION"))
	fmt.Println("Hello World!")
	db := config.MemDatabase(config.AppConfiguration)
	fmt.Println(db)
}
