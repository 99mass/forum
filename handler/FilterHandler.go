package handler

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Filter(db *sql.DB)	http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		_Categorystring := r.Form["category"]
		fmt.Println(_Categorystring)
		date1 := r.FormValue("date1")
		fmt.Println(date1)
	}

}
