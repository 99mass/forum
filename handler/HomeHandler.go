package handler

import (
	"database/sql"
	"forum/controller"
	"forum/helper"
	"forum/middlewares"
	"forum/models"
	"net/http"
)

func Index(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ok, pageError := middlewares.CheckRequest(r, "/", "get")
		if !ok {
			helper.ErrorPage(w, pageError)
			return
		}

		data, err := helper.GetPostForHome(db)
		if err != nil {
			helper.ErrorPage(w, pageError)
			return
		}
		var homeData models.Home
		var sessiondata bool

		sessionID, errsess := helper.SessionHandler(w, r)
		if errsess != nil {

			sessiondata = false
		} else {

			sessiondata = true

			_, errgets := controller.GetSessionByID(db, sessionID)
			if errgets != nil {
				sessiondata = false
			}
		}

		homeData.Session = sessiondata
		homeData.Datas = data

		helper.RenderTemplate(w, "index", "index", homeData)

	}
}
