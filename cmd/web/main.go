package main

import (
	"fmt"
	"log"
	"net/http"
	"webapplication/pkg/config"
	"webapplication/pkg/handlers"
	"webapplication/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.Appconfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port number %s ", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
