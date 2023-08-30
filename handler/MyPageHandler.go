package handler

import (
	"database/sql"
	"forum/controller"
	"forum/helper"
	"forum/models"
	"net/http"

	"github.com/gofrs/uuid"
)

func GetMypage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Check if the user is connected
		var sessiondata bool

		sessionID, errsess := helper.GetSessionRequest(r)
		if errsess != nil {
			sessiondata = false
		} else {

			sessiondata = true

			_, errgets := controller.GetSessionByID(db, sessionID)
			if errgets != nil {
				sessiondata = false
			}

		}
		if sessiondata {
			user, err := controller.GetUserBySessionId(sessionID, db)
			if err != nil {
				controller.DeleteSession(db, sessionID)
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
				helper.RenderTemplate(w, "mypage", "mypages", datas)
			} else {
				PostsDetails, err := helper.GetPostsForOneUser(db, user.ID)
				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
					return
				}
				datas := new(models.DataMypage)
				datas.Datas = PostsDetails

				// dataliked := models.Home{}
				// for _, post := range datas.Datas {
				// 	liked, err := helper.IsPostliked(db, datas.User.ID, post.Posts.ID)
				// 	if err != nil {
				// 		return }
				// 	//Get if disliked
				// 	disliked, errdis := helper.IsPostDisliked(db, datas.User.ID, post.Posts.ID)
				// 	if errdis != nil {
				// 		return
				// 	}
				// 	//fmt.Println(liked)
				// 	if liked {
				// 		post.Liked = true
				// 		dataliked.Datas = append(dataliked.Datas, post)
				// 		continue
				// 	} else if disliked {
				// 		post.Disliked = true
				// 		dataliked.Datas = append(dataliked.Datas, post)
				// 		continue
				// 	} else {
				// 		dataliked.Datas = append(dataliked.Datas, post)
				// 	}
				// }
				// //fmt.Println(dataliked.Datas)
				// fmt.Println(datas.Datas)
				// datas.Datas = dataliked.Datas

				datas.Session = sessiondata
				datas.User = user
				datas.Category = category
				helper.RenderTemplate(w, "mypage", "mypages", datas)
			}

		} else {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

	}
}
