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

func Auth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("auth handler started")

	ok, pageError := middlewares.CheckRequest(r, "/auth", "get")
	if !ok {
		helper.ErrorPage(w, pageError)
		return
	}
	helper.RenderTemplate(w, "auth", "auth", "hello")
	fmt.Println("page rendered")

}

func RegisterHandler(db *sql.DB, templates *template.Template) http.HandlerFunc {
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
