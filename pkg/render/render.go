package render

import (
	"bytes"
	"github/techieasif/bookings/pkg/config"
	"github/techieasif/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//to be used later.
var functions = template.FuncMap{

}

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, page string, td *models.TemplateData) {

	//get the template cache from the application config.
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, hasVal := tc[page]

	if !hasVal {
		log.Fatal("Could not get template from template cache")
	}

	//we are not reading it from disk, that means it has to converted into bytes first.

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)

	if err != nil {
		log.Println("Cannot show template from bytes to browser", err)
	}
}

//CreateTemplateCache creates a template as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		//ts - template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, nil

}
