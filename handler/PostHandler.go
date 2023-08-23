package handler

import (
	"database/sql"
	"forum/controller"
	"forum/helper"
	"forum/models"
	"net/http"

	"github.com/gofrs/uuid"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {

}

func GetOnePost(w http.ResponseWriter, r *http.Request) {

	helper.RenderTemplate(w, "post", "posts", nil)
}
func GetCategorie(w http.ResponseWriter, r *http.Request) {

	helper.RenderTemplate(w, "categorie", "categories", nil)
}

func AddPostHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		session,err := helper.GetSessionRequest(r)
		if err != nil {
			return
		}

		if helper.VerifySession(db, session) {
			postTitle := r.FormValue("")
			postContent := r.FormValue("")
			postCategorystring := r.FormValue("")
			postCategoryuuid,_ := uuid.FromString(postCategorystring)
			user:=controller.GetUserBySessionId(session,db)
			
			post := models.Post{
				UserID: user.ID,
				Title: postTitle,
				Content: postContent,
				CategoryID: postCategoryuuid,
			}
			_,err := controller.CreatePost(db,post)
			if err != nil {
				return
			}
		}
	}
}
