package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

func RenderTemplate(w http.ResponseWriter, tmplName string, tmplDir string, data interface{}) {

	templateCache, err := createTemplateCache(tmplDir)

	if err != nil {
		panic(err)
	}
	// templateCache["home.page.tmpl"]
	tmpl, ok := templateCache[tmplName+".page.tmpl"]

	if !ok {
		http.Error(w, "le template n'existe pas", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, data)
	buffer.WriteTo(w)
}

func createTemplateCache(tmplDir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./template/pages/" + tmplDir + "/*.page.tmpl")
	if err != nil {
		return cache, err
	}
	fmt.Println("not error cache")
	for _, page := range pages {
		//fmt.Println("for pages :", page)
		name := filepath.Base(page)
		//fmt.Println(name)
		tmpl := template.Must(template.ParseFiles(page))

		//fmt.Println(tmpl.Name())
		layouts, err := filepath.Glob("./template/layouts/*.layout.tmpl")
		if err != nil {
			return cache, err
		}
		//fmt.Println("not layout error")
		if len(layouts) > 0 {
			tmpl.ParseGlob("./template/layouts/*.layout.tmpl")
		}
		cache[name] = tmpl
	}
	return cache, nil
}

func RenderError(w http.ResponseWriter, tmplName string, tmplDir string) {
	log.Println("RenderErrorStart")
	templateCache, err := createTemplateCache(tmplDir)

	if err != nil {
		fmt.Println("error detected")
		return
	}
	// templateCache["home.page.tmpl"]
	fmt.Println("RenderErro templateCach")
	tmpl, ok := templateCache[tmplName+".page.tmpl"]

	if !ok {
		RenderTemplate(w, "404", "error", 404)
		//http.Error(w, "le template n'existe pas", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, nil)
	buffer.WriteTo(w)
	log.Println("RenderError end")
}

func ErrorPage(w http.ResponseWriter, i int) error {
	DataError := struct {
		Code    string
		Message string
	}{
		Code:    strconv.Itoa(i),
		Message: http.StatusText(i),
	}
	page, err := template.ParseFiles("template/Error/Errorpage.html")
	if err != nil {
		return err
	}
	w.WriteHeader(i)
	return page.Execute(w, DataError)
}
