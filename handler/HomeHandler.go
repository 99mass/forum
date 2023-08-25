package handler

import (
	"database/sql"
	"fmt"
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
			fmt.Println("err: ", err)
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}
		var homeData models.Home
		var sessiondata bool

		sessionID, errsess := helper.GetSessionRequest(r)
		if errsess != nil {
			sessiondata = false
		} else {

			sessiondata = true

			session, errgets := controller.GetSessionByID(db, sessionID)
			if errgets != nil || &session == nil {
				sessiondata = false
			}
			homeData.User = controller.GetUserBySessionId(sessionID, db)
		}

		category, err := controller.GetAllCategories(db)
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)

		}
		homeData.Session = sessiondata
		homeData.Category = category
		homeData.Datas = data
		if sessiondata {
			helper.UpdateCookieSession(w, sessionID, db)
		}

		helper.RenderTemplate(w, "index", "index", homeData)

	}
}
