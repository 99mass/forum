package handler

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/gofrs/uuid"

	"forum/controller"
	"forum/helper"
	"forum/middlewares"
	"forum/models"
)

func SinginHandler(db *sql.DB) http.HandlerFunc {
	var homeData models.Home
	homeData.Session = false
	homeData.ErrorAuth.EmailError = ""
	homeData.ErrorAuth.GeneralError = ""
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			//Delete the session it was connected
			helper.DeleteSession(w, r)

			ok, pageError := middlewares.CheckRequest(r, "/signin", "get")
			if !ok {
				helper.ErrorPage(w, pageError)
				return
			}
			helper.RenderTemplate(w, "signin", "auth", homeData)

		case http.MethodPost:
			ok, pageError := middlewares.CheckRequest(r, "/signin", "post")
			if !ok {
				helper.ErrorPage(w, pageError)
				return
			}

			datas, err := helper.GetDataTemplate(db, r, false, false, false, true, false)

			if err != nil {
				helper.RenderTemplate(w, "signin", "auth", datas)
				return
			}
			nul := uuid.UUID{}
			if datas.User.ID != nul {
				sess, err := controller.GetSessionIDForUser(db, datas.User.ID)
				if err == nil {
					err := controller.DeleteSession(db, sess)
					if err != nil {
						helper.ErrorPage(w, http.StatusInternalServerError)
						return
					}
				}
				helper.AddSession(w, datas.User.ID, db)
				// Redirect to home
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				helper.RenderTemplate(w, "signin", "auth", datas)
			}
		default:
			helper.ErrorPage(w, http.StatusMethodNotAllowed)
			return
		}
	}
}

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	var homeData models.Home
	homeData.Session = false

	return func(w http.ResponseWriter, r *http.Request) {

		// Handle according to the method
		switch r.Method {
		case http.MethodPost:
			ok, pageError := middlewares.CheckRequest(r, "/register", "post")
			if !ok {
				helper.ErrorPage(w, pageError)
				return
			}
			username := r.FormValue("username")
			username = strings.TrimSpace(username)
			email := r.FormValue("email")
			email = strings.TrimSpace(email)
			password := r.FormValue("password")
			password = strings.TrimSpace(password)
			confirmPassword := r.FormValue("password_validation")
			confirmPassword = strings.TrimSpace(confirmPassword)
			// Hasher le mot de passe
			hashedPassword, _ := helper.HashPassword(password)

			ok, ErrAuth := helper.CheckRegisterFormat(username, email, password, confirmPassword, db)

			if !ok {
				homeData.ErrorAuth = ErrAuth
				helper.RenderTemplate(w, "register", "auth", homeData)
				homeData.ErrorAuth = models.ErrorAuth{}
				return
			}

			user := models.User{
				Username:  username,
				Email:     email,
				Password:  hashedPassword,
				CreatedAt: time.Now(),
			}

			id, _ := controller.CreateUser(db, user)

			// create a session - TODO
			helper.AddSession(w, id, db)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			//helper.RenderTemplate(w, "index", "index", "homedata")
			return

		case http.MethodGet:
			helper.DeleteSession(w, r)
			ok, pageError := middlewares.CheckRequest(r, "/register", "get")
			if !ok {
				helper.ErrorPage(w, pageError)
				return
			}
			helper.RenderTemplate(w, "register", "auth", homeData)
		default:
			helper.ErrorPage(w, http.StatusMethodNotAllowed)
			return
		}
	}
}
