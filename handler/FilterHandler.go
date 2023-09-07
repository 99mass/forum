package handler

import (
	"database/sql"
	"fmt"
	"forum/controller"
	"forum/helper"
	"forum/models"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
)

func Filter(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
			return
		}
		var filterPosts []models.Post
		var category []models.Category
		posts, err := controller.GetAllPosts(db)
		Categorystring := r.Form["category"]
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}

		if Categorystring != nil {

			for _, v := range Categorystring {
				v = strings.TrimSpace(v)
				var cat models.Category
				catuuid, _ := uuid.FromString(v)
				cat.ID = catuuid
				category = append(category, cat)
			}
			postcat, err := controller.GetAllPostCategories(db)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			var idpost []uuid.UUID
			for _, pc := range postcat {
				for _, categ := range category {
					if pc.CategoryID == categ.ID {
						idpost = append(idpost, pc.PostID)
					}
				}
			}
			idpost = RemoveDuplicates(idpost)
			for _, post := range posts {
				for _, postid := range idpost {
					if postid == post.ID {
						filterPosts = append(filterPosts, post)
					}
				}
			}

		}

		if filterPosts == nil {
			filterPosts = posts
		}

		for _, v := range filterPosts {
			fmt.Println(v.Title)
		}

		date1 := r.FormValue("date1")
		fmt.Println(date1)
	}

}


// RemoveDuplicates removes duplicate elements from a slice of uuid.UUID values.
func RemoveDuplicates(input []uuid.UUID) []uuid.UUID {
	unique := make(map[uuid.UUID]bool)
	result := []uuid.UUID{}

	for _, item := range input {
		if !unique[item] {
			unique[item] = true
			result = append(result, item)
		}
	}

	return result
}