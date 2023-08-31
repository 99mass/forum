package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"

	"forum/controller"
	"forum/helper"
	"forum/middlewares"
	"forum/models"
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
			Content = strings.TrimSpace(Content)

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
				homeData.Error = "comments cannot be empty"
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
				homeData, err := helper.GetDataTemplate(db, r, true, false, true, false, true)

				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
					return
				}

				if homeData.Session {
					sessionID, _ := helper.GetSessionRequest(r)
					helper.UpdateCookieSession(w, sessionID, db)
				}
				homeData.Error = errForm.Error()

				helper.RenderTemplate(w, "index", "index", homeData)
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
			fmt.Println(_postCategories)

			user, err := controller.GetUserBySessionId(session, db)
			if err != nil {
				controller.DeleteSession(db, session)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			post := models.Post{
				UserID:     user.ID,
				Title:      postTitle,
				Content:    postContent,
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
			sessiondata := true

			errForm := helper.CheckFormAddPost(r, db)
			if errForm != nil {
				user, err := controller.GetUserBySessionId(session, db)
				if err != nil {
					controller.DeleteSession(db, session)
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
				category, err := controller.GetAllCategories(db)
				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
					return
				}
				CatId := r.FormValue("categorie")
				if CatId != "" {
					CategID, err := uuid.FromString(CatId)
					if err != nil {
						helper.ErrorPage(w, http.StatusBadRequest)
						return
					}
					PostsDetails, err := helper.GetPostsForOneUserAndCategory(db, user.ID, CategID)
					if err != nil {
						helper.ErrorPage(w, http.StatusBadRequest)
					}

					datas := new(models.DataMypage)
					datas.Datas = PostsDetails
					datas.Session = sessiondata
					datas.User = user
					datas.CategoryID = CategID
					datas.Category = category
					datas.Error = errForm.Error()
					helper.RenderTemplate(w, "mypage", "mypages", datas)
				} else {
					PostsDetails, err := helper.GetPostsForOneUser(db, user.ID)
					if err != nil {
						helper.ErrorPage(w, http.StatusInternalServerError)
						return
					}
					datas := new(models.DataMypage)
					datas.Datas = PostsDetails

					datas.Session = sessiondata
					datas.User = user
					datas.Category = category
					datas.Error = errForm.Error()
					helper.RenderTemplate(w, "mypage", "mypages", datas)
				}
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

// Like post
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

// dDislike posts
func DislikePoste(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dislike := models.PostDislike{}

		ok, errorPage := middlewares.CheckRequest(r, "/dislikepost", "post")
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
			dislike.UserID = User.ID
		}

		postID, _ := helper.StringToUuid(r, "post_id")

		dislike.PostID = postID
		_, err := controller.CreatePostDislike(db, dislike)
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}
}

// Like Comments
func LikeComment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		like := models.CommentLike{}

		ok, errorPage := middlewares.CheckRequest(r, "/likecomment", "post")
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

		commentID, _ := helper.StringToUuid(r, "comment_id")

		like.CommentID = commentID
		_, err := controller.CreateCommentLike(db, like)
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}
}

// Dislike comments
func DislikeComment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dislike := models.CommentDislike{}

		ok, errorPage := middlewares.CheckRequest(r, "/dislikecomment", "post")
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
			dislike.UserID = User.ID
		}

		commentID, _ := helper.StringToUuid(r, "post_id")

		dislike.CommentID = commentID
		_, err := controller.CreateCommentDislike(db, dislike)
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}
}
