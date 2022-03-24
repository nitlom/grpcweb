package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/nitlom/webserver/pkg/config"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	// Get the templateCache from the Config

	tc := app.TemplateCache

	t, ok := tc[tmpl]
	if !ok {
		log.Println("Could not fint the template asked for")
	}

	buffer := new(bytes.Buffer)

	err := t.Execute(buffer, nil)
	_, err = buffer.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to the browser", err)
		return
	}
}

// createTemplateCache create a template Cache as a map[string]
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		// fmt.Println("Page is currently:", page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
