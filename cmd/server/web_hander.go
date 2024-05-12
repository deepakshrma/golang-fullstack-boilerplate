package main

import (
	"html/template"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var pathToTemplates = "./templates/"

type TemplateData struct {
	Data  map[string]any
	Error string
	Flash string
}
type staticHandler struct {
	staticPath string
	indexPath  string
}

func (h staticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := filepath.Join(h.staticPath, strings.TrimPrefix(r.URL.Path, "/static"))

	fi, err := os.Stat(p)
	if os.IsNotExist(err) || fi.IsDir() {
		_ = app.render(w, r, "home.gohtml", &TemplateData{})
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.StripPrefix("/static", http.FileServer(http.Dir(h.staticPath))).ServeHTTP(w, r)
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var td = make(map[string]any)

	_ = app.render(w, r, "home.gohtml", &TemplateData{Data: td})

}
func (app *application) render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) error {
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t), path.Join(pathToTemplates, "base.gohtml"))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}
	td.Error = ""
	td.Flash = ""
	err = parsedTemplate.Execute(w, td)
	if err != nil {
		return err
	}
	return nil
}
