package render

import (
	"myapp/pkg/config"
	"fmt"
	"net/http"
	"html/template"
	"path/filepath"
	"log"
	"bytes"
)

var functions = template.FuncMap {}
var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {

	app = a

}

//render html templates
func RenderTemplate(w http.ResponseWriter, tmpl string){

	var tc map[string]*template.Template

	if app.UseCache{
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	//get requested template from cache
	t , ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	buf := new(bytes.Buffer)
	_= t.Execute(buf,nil)

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache()(map[string]*template.Template, error){
	myCache := map[string]*template.Template{}

	//get all of the files named *.pages.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range through all file ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts , err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache , err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache , err
		}

		if len(matches) > 0 {
			ts ,err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache , err
			}
		}

		myCache[name] = ts
	}

	return myCache , nil

}