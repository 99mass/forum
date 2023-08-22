package handler

import (
	"forum/helper"
	"net/http"
)

func GetProfil(w http.ResponseWriter, r *http.Request) {
	helper.RenderTemplate(w, "profil", "profils", nil)
}

func GetMypage(w http.ResponseWriter, r *http.Request) {

	helper.RenderTemplate(w, "mypage", "mypages", nil)
}
