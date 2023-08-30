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
			category, err := controller.GetCategoriesByPost(db, homeData.PostData.Posts.ID)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			homeData.Category = category
			homeData.Datas = posts

			helper.RenderTemplate(w, "post", "posts", homeData)
		case http.MethodPost:
			var comment models.Comment

			postID, errP := helper.StringToUuid(r, "post_id")
			userID, errU := helper.StringToUuid(r, "user_id")
			Content := r.FormValue("content")

			if errP != nil || errU != nil {
				helper.ErrorPage(w, http.StatusBadRequest)
				return
			}
			homeDataSess, err := helper.GetDataTemplate(db, r, true, false, false, false, false)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
			}
			if !homeDataSess.Session {
				homeData, err := helper.GetDataTemplate(db, r, true, true, false, false, false)
				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
				}
				posts, err := helper.GetPostsForOneUser(db, homeData.PostData.User.ID)
				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
					return
				}
				category, err := controller.GetCategoriesByPost(db, homeData.PostData.Posts.ID)
				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
					return
				}
				homeData.Category = category
				homeData.Datas = posts
				helper.RenderTemplate(w, "post", "posts", homeData)
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

			_, erro := controller.CreateComment(db, comment)
			if erro != nil {
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
			category, err := controller.GetCategoriesByPost(db, homeData.PostData.Posts.ID)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			homeData.Category = category
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
				//helper.ErrorPage(w, http.StatusBadRequest)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			postTitle := r.FormValue("title")
			postContent := r.FormValue("content")
			_postCategorystring := r.Form["category"]
			if _postCategorystring == nil || postTitle == "" || postContent == "" {
				homeData, err := helper.GetDataTemplate(db, r, true, false, true, false, true)

				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
					return
				}

				if homeData.Session {
					sessionID, _ := helper.GetSessionRequest(r)
					helper.UpdateCookieSession(w, sessionID, db)
				}
				homeData.Error = "please complete all fields"

				helper.RenderTemplate(w, "index", "index", homeData)
				return
			}
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

func LikePoste(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		like := models.PostLike{}

		ok, errorPage := middlewares.CheckRequest(r, "/likepost", "post")
		if !ok {
			helper.ErrorPage(w, errorPage)
			return
		}

		//check the session and get the user
		sessionID, errsess := helper.GetSessionRequest(r)
		if errsess != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {

			session, errgets := controller.GetSessionByID(db, sessionID)
			if errgets != nil || &session == nil {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			User, errgetu := controller.GetUserBySessionId(sessionID, db)
			if errgetu != nil {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			like.UserID = User.ID
		}

		postID, _ := helper.StringToUuid(r, "post_id")

		like.PostID = postID
		_, err := controller.CreatePostLike(db, like)
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}
}
