package handler

import (
	"database/sql"
	"forum/controller"
	"forum/helper"
	"forum/models"
	"net/http"
)

func GetProfil(w http.ResponseWriter, r *http.Request) {
	helper.RenderTemplate(w, "profil", "profils", nil)
}

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
			user := controller.GetUserBySessionId(sessionID, db)
			PostsDetails, err := helper.GetPostForMyPage(db, user.ID)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			category, err := controller.GetAllCategories(db)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return 
			}
			datas := new(models.DataMypage)
			datas.Datas = PostsDetails
			datas.Session = sessiondata
			datas.User = user
			datas.Category = category
			helper.RenderTemplate(w, "mypage", "mypages", datas)
		} else {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

	}
}
