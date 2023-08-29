package handler

import (
	"database/sql"
	"forum/controller"
	"forum/helper"
	"forum/middlewares"
	"forum/models"
	"net/http"

	"github.com/gofrs/uuid"
)

func GetOnePost(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			homeData, err := helper.GetDataTemplate(db, r, true, true, false, false, false)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			posts, err := helper.GetPostsForOneUser(db, homeData.PostData.User.ID)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			homeData.Datas = posts

			helper.RenderTemplate(w, "post", "posts", homeData)
		case http.MethodPost:
			var comment models.Comment

			postID, errP := helper.StringToUuid(r, "post_id")
			userID, errU := helper.StringToUuid(r, "user_id")
			Content := r.FormValue("content")

			if errP != nil || errU != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			if Content == "" {
				homeData, err := helper.GetDataTemplate(db, r, true, true, false, false, false)
				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
				}
				posts, err := helper.GetPostsForOneUser(db, homeData.PostData.User.ID)
				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
					return
				}
				homeData.Datas = posts
				helper.RenderTemplate(w, "post", "posts", homeData)
				return
			}
			comment.PostID = postID
			comment.UserID = userID
			comment.Content = Content

			_, err := controller.CreateComment(db, comment)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			homeData, err := helper.GetDataTemplate(db, r, true, true, true, false, false)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
			}
			posts, err := helper.GetPostsForOneUser(db, homeData.User.ID)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			homeData.Datas = posts
			helper.RenderTemplate(w, "post", "posts", homeData)

		}

	}
}

func AddPostHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ok, errorPage := middlewares.CheckRequest(r, "/addpost", "post")
		if !ok {
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
				helper.ErrorPage(w, http.StatusBadRequest)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			postTitle := r.FormValue("title")
			postContent := r.FormValue("content")
			_postCategorystring := r.Form["category"]
			// var _postCategoryuuid []uuid.UUID
			var _postCategories []models.Category
			// for _, v := range _postCategorystring {
			// 	catuuid, _ := uuid.FromString(v)
			// 	_postCategoryuuid = append(_postCategoryuuid, catuuid)
			// }
			for _, v := range _postCategorystring {
				var cat models.Category
				catuuid, _ := uuid.FromString(v)
				cat.ID = catuuid
				_postCategories = append(_postCategories, cat)
			}

			user, err := controller.GetUserBySessionId(session, db)
			if err != nil {
				controller.DeleteSession(db, session)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			post := models.Post{
				UserID:  user.ID,
				Title:   postTitle,
				Content: postContent,
				// CategoryID: _postCategoryuuid,
				Categories: _postCategories,
			}
			_, err = controller.CreatePost(db, post)
			if err != nil {
				return
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	}
}

func AddPostHandlerForMyPage(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ok, errorPage := middlewares.CheckRequest(r, "/addpostmypage", "post")
		if !ok {
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
				helper.ErrorPage(w, http.StatusBadRequest)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			postTitle := r.FormValue("title")
			postContent := r.FormValue("content")
			_postCategorystring := r.Form["category"]
			// var _postCategoryuuid []uuid.UUID
			var _postCategories []models.Category
			// for _, v := range _postCategorystring {
			// 	catuuid, _ := uuid.FromString(v)
			// 	_postCategoryuuid = append(_postCategoryuuid, catuuid)
			// }
			for _, v := range _postCategorystring {
				var cat models.Category
				catuuid, _ := uuid.FromString(v)
				cat.ID = catuuid
				_postCategories = append(_postCategories, cat)
			}

			user, err := controller.GetUserBySessionId(session, db)
			if err != nil {
				controller.DeleteSession(db, session)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			post := models.Post{
				UserID:  user.ID,
				Title:   postTitle,
				Content: postContent,
				// CategoryID: _postCategoryuuid,
				Categories: _postCategories,
			}
			_, err = controller.CreatePost(db, post)
			if err != nil {
				return
			}
			http.Redirect(w, r, "/mypage", http.StatusSeeOther)
		}

	}
}
