package middlewares

import (
	"net/http"
	"strings"
)

func CheckRequest(r *http.Request, path, methode string) (bool, int) {
	if strings.ToLower(r.Method) == methode && r.URL.Path == path{
		return true, 0
	} else if !Getmethode(r,methode){
		return false, 405
	} else {
		return false, 404
	}
}

func Getmethode(r *http.Request, methode string) (bool){
	if strings.ToLower(r.Method) != methode{
		return false
	}
	return true
}