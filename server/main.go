package main

import (
	"boilerplate/env"
	"boilerplate/template"
	"fmt"
	"os"
)

func init() {
	env.LoadEnvs()
	template.LoadTemplates()
}

func main() {

	template.Templates["version.tmpl"].Execute(os.Stdout, os.Getenv("APP_VERSION"))
	fmt.Println("Hello World!")
}
