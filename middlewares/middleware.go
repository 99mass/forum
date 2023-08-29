package middlewares

import (
	"forum/helper"
	"net/http"
	"strings"
)

func CheckRequest(r *http.Request, path, methode string) (bool, int) {
	//helper.Debug(r.Method + " " + methode + " " + r.URL.Path + " " + path)
	//routes:= []string{}
	if strings.ToLower(r.Method) == methode && r.URL.Path == path {
		helper.Debug("Request validated")
		return true, 0
	} else if strings.ToLower(r.Method) != methode {
		helper.Debug("Request Methode doesn't match")
		return false, 405
	} else {
		helper.Debug("Requset page not found")
		
		return false, 404
	}
}
