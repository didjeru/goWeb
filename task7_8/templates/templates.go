package templates

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"sync"
)

// Templates all templates
type Templates struct {
	templates map[string]*template.Template
	mux       sync.Mutex
}

// GetTemplateByName get template by name
func (tempts *Templates) GetTemplateByName(name string) *template.Template {
	tempts.mux.Lock()
	tmpl := tempts.templates[name]
	tempts.mux.Unlock()
	return tmpl
}

// Init init templates
func Init() Templates {
	tempts := make(map[string]*template.Template)
	files := getFilesFromDir(getTemplatesPath())

	for _, f := range files {
		tpl, err := template.ParseFiles(path.Join(getTemplatesPath(), f.Name()))
		if err != nil {
			log.Fatal(err)
		}

		tempts[filenameWithoutExtension(f.Name())] = tpl
	}

	return Templates{templates: tempts}
}

func filenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func getFilesFromDir(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func getTemplatesPath() string {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return path.Join(dir, "templates", "html")
}
