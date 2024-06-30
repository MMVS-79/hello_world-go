package main

import (
	"fmt"
	"net/http"
	"html/template"
)

//render html templates
func renderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate,_:=template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w,nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}