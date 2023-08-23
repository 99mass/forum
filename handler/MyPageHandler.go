package handler

import (
	"forum/helper"
	"net/http"
)

func MyPageHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the user is connected

	//
	helper.RenderTemplate(w,"index","index","data")
}
