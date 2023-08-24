package handler

import (
	"database/sql"
	"forum/controller"
	"forum/helper"
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
			helper.RenderTemplate(w, "mypage", "mypages", nil)
		} else {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

	}
}
