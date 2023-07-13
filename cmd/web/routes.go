package main

import (
	"net/http"
	"webapplication/pkg/config"
	"webapplication/pkg/handlers"

	"github.com/bmizerany/pat"
)

func routes(app *config.Appconfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
