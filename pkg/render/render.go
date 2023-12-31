package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"webapplication/pkg/config"

	"webapplication/pkg/models"
)

var functions = template.FuncMap{}
var app *config.Appconfig

// NewTemplate sets the configuration for the template package
func NewTemplates(a *config.Appconfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// renderTemplate uses html template
func RenderTemplate(w http.ResponseWriter, html string, td *models.TemplateData) {
	//get templateCache from appConfig
	var tc map[string]*template.Template
	if app.UseCache {

		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[html]
	if !ok {
		log.Fatal("could not get template from template cache")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)
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
