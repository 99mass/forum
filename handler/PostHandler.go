package handler

import (
	"database/sql"
	"fmt"
	"forum/controller"
	"forum/helper"
	"forum/middlewares"
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
		ok, errorPage := middlewares.CheckRequest(r, "/addpost", "post")
		if !ok {
			helper.Debug("Checkrequest addpost failled")
			helper.ErrorPage(w, errorPage)
			return
		}
		session, err := helper.GetSessionRequest(r)
		if err != nil {
			return
		}

		if helper.VerifySession(db, session) {

			errForm := helper.CheckFormAddPost(r, db)
			if errForm != nil {
				helper.Debug(errForm.Error())
				helper.ErrorPage(w, http.StatusBadRequest)
				http.Redirect(w, r, "/",http.StatusSeeOther)
				return
			}
			postTitle := r.FormValue("title")
			postContent := r.FormValue("content")
			_postCategorystring := r.Form["category"]
			fmt.Println("hello:",_postCategorystring)
			var _postCategoryuuid []uuid.UUID
			for _, v := range _postCategorystring {
				catuuid, _ := uuid.FromString(v)
				_postCategoryuuid = append(_postCategoryuuid, catuuid)
			}

			user := controller.GetUserBySessionId(session, db)

			post := models.Post{
				UserID:     user.ID,
				Title:      postTitle,
				Content:    postContent,
				CategoryID: _postCategoryuuid,
			}
			_,err := controller.CreatePost(db, post)
			if err != nil {
				fmt.Println(err," pos no cre")
				return
			}
			fmt.Println("hello")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		
	}
}
