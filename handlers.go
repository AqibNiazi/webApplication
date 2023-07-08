package main

import (
	"net/http"
)

// Home is the home page handler
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

// About is the about page handler
func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}
