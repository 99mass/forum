package handler

import (
	"fmt"
	"forum/helper"
	"forum/middlewares"
	"net/http"
)



func Auth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("auth handler started")

	
		ok, pageError:= middlewares.CheckRequest(r, "/auth", "get")
		if !ok {
			helper.ErrorPage(w, pageError)
			return
		}
		helper.RenderTemplate(w, "auth", "auth", "hello")
		fmt.Println("page rendered")
	
}