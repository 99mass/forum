package handler

import (
	"database/sql"
	"forum/controller"
	"forum/helper"
	"forum/middlewares"
	"forum/models"
	"net/http"
	"time"
)



func SinginHandler(db *sql.DB) http.HandlerFunc {
	var homeData models.Home
	homeData.Session = false

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
			email := r.FormValue("email")
			password := r.FormValue("motdepasse")

			//Check if the error has to be handled
			userID, toConnect := helper.VerifUser(db, email, password)

			if toConnect {
				// Create a session
				helper.AddSession(w, userID, db)
				// Redirect to home
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				helper.RenderTemplate(w, "signin", "auth", homeData)
			}

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
			helper.Debug("envoie du formulaire")
			ok, pageError := middlewares.CheckRequest(r, "/register", "post")
			if !ok {
				helper.ErrorPage(w, pageError)
				return
			}
			username := r.FormValue("username")
			email := r.FormValue("email")
			password := r.FormValue("password")
			confirmPassword := r.FormValue("password_validation")
			// Hasher le mot de passe
			hashedPassword, _ := helper.HashPassword(password)

			ok, ErrAuth := helper.CheckRegisterFormat(username, email, password, confirmPassword, db)

			if !ok {
				homeData.ErrorAuth = ErrAuth
				helper.RenderTemplate(w, "register", "auth", homeData)
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
			//fmt.Println("affichage du formulaire d'enregistrement")
			ok, pageError := middlewares.CheckRequest(r, "/register", "get")
			if !ok {
				helper.ErrorPage(w, pageError)
				return
			}
			helper.RenderTemplate(w, "register", "auth", homeData)
		default:
			helper.ErrorPage(w, 404)
			return
		}
	}
}
