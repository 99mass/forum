package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"

	// "forum/handler"
	"forum/controller"
	"forum/helper"
	"forum/models"
	"forum/routes"
)

var PORT = ":8080"

// Structure pour stocker les informations de l'utilisateur
type UserData struct {
	Username string
	Password string
}

// Structure pour stocker les données du formulaire de mise à jour
type UpdateFormData struct {
	OldPassword string
	Username    string
	NewPassword string
}

var user UserData // Pour stocker les informations de l'utilisateur

func main() {
	db, _ := helper.CreateDatabase()

	err := helper.CreateTables(db)
	if err != nil {
		fmt.Println(err)
	}
	user := new(models.User)
	user.Username = "mouhamed"
	user.Email = "mouha@gmail.com"
	user.Password = "password"
	id, _ := controller.CreateUser(db, *user)
	fmt.Println(id)
	fmt.Println(user.ID)

	args := os.Args[1:]
	if len(args) > 0 {
		tab := strings.Split(args[0], "=")
		if len(tab) == 2 && helper.IsInt(tab[1]) && tab[1] != "" && tab[0] == "--port" {
			t, _ := strconv.Atoi(tab[1])
			if t >= 1024 && t <= 65535 {
				PORT = ":" + tab[1]
			}
		}
	}
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	routes.Route()
	fmt.Println("Listening in http://localhost" + PORT)

	http.ListenAndServe(PORT, nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl := template.Must(template.ParseFiles("./template/pages/index.html"))
			tmpl.Execute(w, nil)
		}
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()

			formData := UpdateFormData{
				OldPassword: r.FormValue("old_password"),
				Username:    r.FormValue("username"),
				NewPassword: r.FormValue("new_password"),
			}

			// Vérification de l'ancien mot de passe
			if formData.OldPassword != user.Password {
				http.Error(w, "Ancien mot de passe incorrect", http.StatusUnauthorized)
				return
			}

			// Mise à jour des informations de l'utilisateur (simulé ici)
			user.Username = formData.Username
			user.Password = formData.NewPassword

			// Une fois la mise à jour terminée, vous pouvez rediriger l'utilisateur vers une page de succès ou une autre page
			http.Redirect(w, r, "/success", http.StatusSeeOther)
		}
	})

	http.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
		// Page de succès
		w.Write([]byte("Mise à jour réussie !"))
	})

}


