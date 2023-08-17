package handler

import (
	"forum/helper"
	"forum/middlewares"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	ok, pageError := middlewares.CheckRequest(r, "/", "get")
	if !ok {
		helper.ErrorPage(w, pageError)
		return
	}
	helper.RenderTemplate(w, "index", "index", "hello")

}
