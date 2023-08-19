package handler

import (
	"database/sql"
	"fmt"
	"forum/controller"
	"forum/helper"
	"forum/middlewares"
	"forum/models"
	"html/template"
	"net/http"
	"time"
)

func SinginHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Println("auth signin handler started")
			ok, pageError := middlewares.CheckRequest(r, "/signin", "get")
			if !ok {
				helper.ErrorPage(w, pageError)
				return
			}
			helper.RenderTemplate(w, "signin", "auth", "")
			fmt.Println("page rendered")
		case http.MethodPost:
			username := r.FormValue("user_name")
			email := r.FormValue("email")
			password := r.FormValue("password")
			//Check if the error has to be handled
			hashedPassword, _ := helper.HashPassword(password)

			user := models.User{
				Username:  username,
				Email:     email,
				Password:  hashedPassword,
				CreatedAt: time.Now(),
			}
			if user.ConnectUser(db) {
				// Create a session

				// Redirect to home
				helper.RenderTemplate(w, "index", "index", user)
			} else {
				helper.RenderTemplate(w, "signin", "auth", user)
			}

		}

	}
}

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	fmt.Println("Register handler")
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPost:
			fmt.Println("envoie du formulaire")
			ok, pageError := middlewares.CheckRequest(r, "register", "post")
			if !ok {
				helper.ErrorPage(w, pageError)
				return
			}
			username := r.FormValue("user_name")
			email := r.FormValue("email")
			password := r.FormValue("password")
			confirmPassword := r.FormValue("password_validation")
			// Hasher le mot de passe
			hashedPassword, _ := helper.HashPassword(password)

			if !helper.ConfirmPasswordsMatch(password, confirmPassword) {
				//http.Error(w, "Les mots de passe ne correspondent pas", http.StatusBadRequest)
				fmt.Println("Les mots de passe ne sont pas conformes")
				helper.RenderTemplate(w, "register", "auth", "error")
				return
			}
			fmt.Println("password matches")

			user := models.User{
				Username:  username,
				Email:     email,
				Password:  hashedPassword,
				CreatedAt: time.Now(),
			}

			_, err := user.Register(db)
			if err != nil {
				//http.Error(w, err.Error(), http.StatusBadRequest)
				fmt.Println(err)
				helper.RenderTemplate(w, "register", "auth", "error")
				return
			}

			// create a session

			//Redirect to home page
			fmt.Println("enregistrement réussi")
			helper.RenderTemplate(w, "index", "index", "homedata")

		case http.MethodGet:
			fmt.Println("affichage du formulaire d'enregistrement")
			ok, pageError := middlewares.CheckRequest(r, "/register", "get")
			if !ok {
				helper.ErrorPage(w, pageError)
				return
			}
			helper.RenderTemplate(w, "register", "auth", "")
		default:
			helper.ErrorPage(w, 404)
			return
		}
	}
}

func RegisterHandlers(db *sql.DB, templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			username := r.FormValue("user_name")
			email := r.FormValue("email")
			password := r.FormValue("password")
			confirmPassword := r.FormValue("password_validation")

			if !helper.ConfirmPasswordsMatch(password, confirmPassword) {
				http.Error(w, "Les mots de passe ne correspondent pas", http.StatusBadRequest)
				return
			}

			// Vérifier si le nom d'utilisateur ou l'adresse e-mail est déjà pris
			duplicate, err := controller.IsDuplicateUsernameOrEmail(db, username, email)
			if err != nil {
				http.Error(w, "Erreur lors de la vérification du duplicata", http.StatusInternalServerError)
				return
			}
			if duplicate {
				http.Error(w, "Nom d'utilisateur ou adresse e-mail déjà pris", http.StatusBadRequest)
				return
			}

			// Hasher le mot de passe
			hashedPassword, err := helper.HashPassword(password)
			if err != nil {
				http.Error(w, "Erreur lors du hachage du mot de passe", http.StatusInternalServerError)
				return
			}

			// Enregistrer l'utilisateur dans la base de données
			user := models.User{
				Username:  username,
				Email:     email,
				Password:  hashedPassword,
				CreatedAt: time.Now(),
			}
			_, err = controller.CreateUser(db, user)
			if err != nil {
				http.Error(w, "Erreur lors de l'enregistrement de l'utilisateur", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Afficher le formulaire d'inscription
		templates.ExecuteTemplate(w, "template/pages/auth/auth.page.tmpl", nil)
	}
}
