package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

//Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("This is the home page"))
}

//About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is %d", sum))
}

//Divide is the divide page handler
func Divide(w http.ResponseWriter, r *http.Request) {
	x := 100.0
	y := 0.0
	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot devide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", x, y, f))
}


//divideValues  
func divideValues(x, y float64) (float64, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}

//addValue adds two integers and return the sum
func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
