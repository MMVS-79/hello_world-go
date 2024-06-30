package main

import (
	"net/http"
	"fmt"
	//"errors"
)

const portNum = ":8080"


func main() {


	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNum))
	_ = http.ListenAndServe(portNum, nil)

}