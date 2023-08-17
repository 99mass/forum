package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	// "forum/handler"
	"forum/controller"
	"forum/helper"
	"forum/models"
	"forum/routes"
)

var PORT = ":8080"

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
}
