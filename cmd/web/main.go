package main

import (
	"net/http"
	"fmt"
	"myapp/pkg/handlers"
	//"errors"
)

const portNum = ":8080"


func main() {


	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNum))
	_ = http.ListenAndServe(portNum, nil)

}