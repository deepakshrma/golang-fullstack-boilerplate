package templates

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"webapp/pkg/config"
)

var Templates = map[string]*template.Template{}

func LoadTemplates() {
	templatesPath := filepath.Join(config.AppWd, "templates")
	files, _ := os.ReadDir(templatesPath)
	for _, file := range files {
		if file.IsDir() && strings.HasSuffix(file.Name(), ".tmpl") {
			continue
		}
		filesStr := filepath.Join(templatesPath, file.Name())
		tmpl := template.Must(template.New(file.Name()).ParseFiles(filesStr))
		Templates[file.Name()] = tmpl
	}
}
