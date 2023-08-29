package handler

import (
	"database/sql"
	"fmt"
	"forum/controller"
	"forum/helper"
	"net/http"

	"github.com/gofrs/uuid"
)

func GetPostCategory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categoryID, err := helper.StringToUuid(r, "categorie")
		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
			return
		}
		if VerifCategory(db, categoryID) {
			
		}

		fmt.Println(categoryID)
	}
}

func VerifCategory(db *sql.DB, Id uuid.UUID) bool {
	category, err := controller.GetCategoryByID(db, Id)
	if err != nil {
		return false
	}
	if &category == nil {
		return false
	}
	return true
}
