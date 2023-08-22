package handler

import (
	"forum/helper"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {

}

func GetOnePost(w http.ResponseWriter, r *http.Request) {

	helper.RenderTemplate(w, "post", "posts", nil)
}
func GetCategorie(w http.ResponseWriter, r *http.Request) {

	helper.RenderTemplate(w, "categorie", "categories", nil)
}
