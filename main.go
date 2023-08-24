package main

import (
	"fmt"
	"net/http"

	// "forum/handler"

	"forum/helper"
	"forum/models"
	"forum/routes"
)

var PORT = ":8080"
var Category = []models.Category{}

//var CategoryName = []string{"Education", "Sport", "Art", "Culture", "Religion"}

func main() {

	// Openning the database
	db, _ := helper.CreateDatabase()

	// Create tables
	err := helper.CreateTables(db)
	if err != nil {
		fmt.Println(err)
	}

	// user := new(models.User)
	// user.Username = "mouhamed"
	// user.Email = "mouha@gmail.com"
	// user.Password = "password"
	// id, _ := controller.CreateUser(db, *user)
	// user.ID = id
	// user2, _ := controller.GetAllUsers(db)
	// for _, v := range user2 {
	// 	fmt.Println("ID:",v.ID)
	// 	fmt.Println("username:",v.Username)
	// 	fmt.Println("Email:",v.Email)
	// 	fmt.Println("password:",v.Password)
	// 	fmt.Println("Created_at:",v.CreatedAt)
	// }

	// Allowing the client to chose the PORT server listenning
	// args := os.Args[1:]
	// if len(args) > 0 {
	// 	tab := strings.Split(args[0], "=")
	// 	if len(tab) == 2 && helper.IsInt(tab[1]) && tab[1] != "" && tab[0] == "--port" {
	// 		t, _ := strconv.Atoi(tab[1])
	// 		if t >= 1024 && t <= 65535 {
	// 			PORT = ":" + tab[1]
	// 		}
	// 	}
	// }
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Run Handlers
	routes.Route(db)

	fmt.Println("Listening in http://localhost" + PORT)
	// for _, v := range CategoryName {
	// 	cat := models.Category{
	// 		NameCategory: v,
	// 	}
	// 	controller.CreateCategory(db, cat)
	// }

	http.ListenAndServe(PORT, nil)
}
