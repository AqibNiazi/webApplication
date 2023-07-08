package main

import (
	"fmt"
	"net/http"
)

const portNumber string = ":8080"

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	fmt.Println(fmt.Sprintf("Starting application on port number %s ", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
