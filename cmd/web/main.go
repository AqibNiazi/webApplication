package main

import (
	"fmt"
	"net/http"
	"webapplication/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port number %s ", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
