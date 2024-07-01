package render

import (
	"fmt"
	"net/http"
	"html/template"
)

//render html templates
func RenderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate,_:=template.ParseFiles("./templates/" + tmpl , "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w,nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}