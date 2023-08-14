package helper

import (
	"html/template"
	"net/http"
	"strconv"
)

func RenderTemplate(w http.ResponseWriter, s string, Data interface{}) error {
	page, err := template.ParseFiles("template/" + s + ".html")
	if err != nil {
		return err
	}
	return page.Execute(w, Data)
}
func RenderTemplateWithLoyout(w http.ResponseWriter, s string, Data interface{}) error {
	page, err := template.ParseFiles("template/layouts/header.layout.tmpl","template/" + s + ".html")
	if err != nil {
		return err
	}
	
	return page.ExecuteTemplate(w,"layout" ,Data)
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
