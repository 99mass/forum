package handler

import (
	"forum/helper"
	"net/http"
)


func SignOutHandler(w http.ResponseWriter, r *http.Request) {

	helper.DeleteSession(w,r)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}