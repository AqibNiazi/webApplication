package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber string = ":8080"

// Home is the home page handler
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page")
}

// About is the about page handler
func about(w http.ResponseWriter, r *http.Request) {
	sum := addValue(3, 4)
	fmt.Fprintf(w, fmt.Sprintf("This is about page and 3 + 4 is %d ", sum))
}

// AddValues add two values and return the sum
func addValue(x, y int) int {
	return x + y
}
func divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValue(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by zero")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}
func divideValue(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/divide", divide)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello world")
	// 	if err != nil {
	// 		fmt.Println("application showing error", err)
	// 	}

	// 	fmt.Printf("number of bytes written %d \n", n)
	// })
	fmt.Println(fmt.Sprintf("Starting application on port number %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
