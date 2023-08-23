package routes

import (
	"database/sql"
	"forum/handler"
	"net/http"
)

func Route(db *sql.DB) {
	http.HandleFunc("/", handler.Index(db))
	http.HandleFunc("/signin", handler.SinginHandler(db))
	http.HandleFunc("/register", handler.RegisterHandler(db))
	http.HandleFunc("/mypage", handler.GetMypage)
	http.HandleFunc("/posts", handler.GetPosts)
	http.HandleFunc("/post/", handler.GetOnePost)
	http.HandleFunc("/categorie/", handler.GetCategorie)
	http.HandleFunc("/profil", handler.GetProfil)
	http.HandleFunc("/signout", handler.SignOutHandler)


}
