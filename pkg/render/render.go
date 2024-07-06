package render

import (
	//"fmt"
	"net/http"
	"html/template"
	"path/filepath"
	"log"
	"bytes"
)

//render html templates
func RenderTemplate(w http.ResponseWriter, tmpl string){
	//create a template cache
	tc , err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	//get requested template from cache
	t , ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf,nil)
	if err != nil {
		log.Println(err)
	}
	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache()(map[string]*template.Template, error){
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