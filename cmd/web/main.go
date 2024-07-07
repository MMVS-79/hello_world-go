package main

import (
	"net/http"
	"fmt"
	"myapp/pkg/handlers"
	"myapp/pkg/config"
	"myapp/pkg/render"
	"log"
	//"errors"
)

const portNum = ":8080"


func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)


	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNum))
	_ = http.ListenAndServe(portNum, nil)

}