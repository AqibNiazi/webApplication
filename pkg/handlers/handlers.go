package handlers

import (
	"net/http"
	"webapplication/pkg/config"
	"webapplication/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.Appconfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
