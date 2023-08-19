package routes

import (
	"database/sql"
	"forum/handler"
	"net/http"
)

func Route(db *sql.DB) {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/signin", handler.SinginHandler(db))
	http.HandleFunc("/register", handler.RegisterHandler(db))
	http.HandleFunc("/posts", handler.GetPosts)
	http.HandleFunc("/post/", handler.GetOnePost)
	http.HandleFunc("/profil/", handler.GetProfil)

}
