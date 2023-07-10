package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"webapplication/pkg/config"
)

var functions = template.FuncMap{}
var app *config.Appconfig

// NewTemplate sets the configuration for the template package
func NewTemplates(a *config.Appconfig) {
	app = a
}

// renderTemplate uses html template
func RenderTemplate(w http.ResponseWriter, html string) {
	//get templateCache from appConfig
	tc := app.TemplateCache

	t, ok := tc[html]
	if !ok {
		log.Fatal("could not get template from template cache")
	}
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error occured while writing to the browser", err)
	}

}

// CreateTemplateCache
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		match, err := filepath.Glob("./templates/*layout.html")
		if err != nil {
			return myCache, err
		}
		if len(match) > 0 {
			ts, err := ts.ParseGlob("./templates/*layout.html")
			if err != nil {
				return myCache, err
			}
			myCache[name] = ts
		}
	}
	return myCache, nil
}
