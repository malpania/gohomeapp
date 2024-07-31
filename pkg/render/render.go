package render

import (
	"fmt"
	"github.com/malpania/beerproj/pkg/config"
	"github.com/malpania/beerproj/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func InitializeApp(AppConfig *config.AppConfig) {
	app = AppConfig
}

func RenderTemplateOld(writer http.ResponseWriter, templateFile string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+templateFile, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(writer, nil)
	if err != nil {
		fmt.Println("Error parsing tempalte : ", err)
		return
	}
}
func addDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

func RenderTemplate(writer http.ResponseWriter, templateFile string, templateData *models.TemplateData) {
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		_ = createTemplateCache(templateFile)
	}
	t, ok := tc[templateFile]
	if !ok {
		log.Fatal("Could not get Template file not found")
	}
	templateData = addDefaultData(templateData)

	err := t.Execute(writer, templateData)
	if err != nil {
		fmt.Println("Error parsing template : ", err)
		return
	}

}

func ReadFolderCache() (map[string]*template.Template, error) {

	myCache := make(map[string]*template.Template)

	pages, err := filepath.Glob("templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		fileName := filepath.Base(page)
		ts, err := template.New(fileName).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}

		}
		myCache[fileName] = ts
	}
	return myCache, nil
}

var tc = make(map[string]*template.Template)

func RenderTemplateOldTest(writer http.ResponseWriter, templateFile string) {

	var tmpl *template.Template
	var err error

	if tc[templateFile] == nil {
		err = createTemplateCache(templateFile)
		if err != nil {
			fmt.Println("Error parsing tempalte : ", err)
			return
		}
	} else {
		tmpl = tc[templateFile]
		err = tmpl.Execute(writer, nil)
	}

}

func createTemplateCache(t string) error {

	templates := []string{
		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		fmt.Println("Error parsing tempalte : ", err)
		return err
	}
	tc[t] = tmpl
	return nil
}
