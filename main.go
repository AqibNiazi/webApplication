package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber string = ":8080"

// Home is the home page handler
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

// About is the about page handler
func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}
func renderTemplate(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error occured while parsing", err)
		return
	}

}
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	fmt.Println(fmt.Sprintf("Starting application on port number %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
