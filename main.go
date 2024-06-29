package main

import (
	"net/http"
	"fmt"
	"errors"
)

const portNum = ":8080"

//Home is the home page handler 
func Home(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "This is the home page")

}


// About is the about page handler
func About(w http.ResponseWriter, r *http.Request){

	sum := addValues(2,2)
	_,_=fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is %d", sum))

}

//addVlaues addes two integers and returns the sum 
func addValues(x,y int) int {

	return x + y 
}

func Divide(w http.ResponseWriter, r *http.Request){

	f, err := divideValues(100.0, 0)
	if err != nil {

		fmt.Fprintf(w, "Cannot divide by 0")
		return

	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))

}

func divideValues(x, y float32)(float32, error){

	if y<=0 {

		err := errors.New("Cannot divide by zero")
		return 0, err
	}

	result := x/y
	return result, nil
}

func main() {


	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNum))
	_=http.ListenAndServe(portNum, nil)

}