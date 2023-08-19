package routes

import (
	"forum/handler"
	"net/http"
)

func Route() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/auth", handler.Auth)
	http.HandleFunc("/posts", handler.GetPosts)
	http.HandleFunc("/post/", handler.GetOnePost)
	http.HandleFunc("/profil/", handler.GetProfil)

}
