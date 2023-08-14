package main

import (
	"fmt"
	// "forum/handler"
	"forum/helper"
	"forum/routes"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var PORT = ":8080"

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		tab := strings.Split(args[0], "=")
		if len(tab) ==2 && helper.IsInt(tab[1]) && tab[1] != "" && tab[0]=="--port" {
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