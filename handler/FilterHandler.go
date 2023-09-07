package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofrs/uuid"

	"forum/controller"
	"forum/helper"
	"forum/models"
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

		

		date1 := r.FormValue("date1")
		date2 := r.FormValue("date2")
		likemi := r.FormValue("likemin")
		likema:= r.FormValue("likemax")
		likemin,err := strconv.Atoi(likemi)
		if err != nil {
			helper.ErrorPage(w,http.StatusBadRequest)
			return
		}
		likemax,err := strconv.Atoi(likema)
		if err != nil {
			helper.ErrorPage(w,http.StatusBadRequest)
			return
		}


		filterPosts, err = GetFilteredPosts(db, filterPosts, date1, date2)
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}
		Posts,err := helper.GetPostForFilter(db,filterPosts)
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}
		var PostsFiltered []models.HomeDataPost
		for _, v := range Posts {
			if v.PostLike >= likemin && v.PostLike <= likemax{
				PostsFiltered = append(PostsFiltered, v)
			}
		}
		for _,v := range PostsFiltered{
			fmt.Println(v.PostLike)
		}
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

func GetFilteredPosts(db *sql.DB, posts []models.Post, minDate, maxDate string) ([]models.Post, error) {
	var filteredPosts []models.Post

	for _, post := range posts {

		createdAt, err := time.Parse("2006-01-02 15:04:05", post.CreatedAt)
		if err != nil {
			return nil, err
		}

		minDateTime, err := time.Parse("2006-01-02", minDate)
		if err != nil {
			return nil, err
		}

		maxDateTime, err := time.Parse("2006-01-02", maxDate)
		if err != nil {
			return nil, err
		}

		// Si aucune heure n'est fournie, ajustez les heures à minuit et 23:59:59
		minDateTime = minDateTime.Add(time.Hour * time.Duration(0))
		maxDateTime = maxDateTime.Add(time.Hour*time.Duration(23) + time.Minute*time.Duration(59) + time.Second*time.Duration(59))

		if createdAt.After(minDateTime) && createdAt.Before(maxDateTime) {
			filteredPosts = append(filteredPosts, post)
		}
	}

	return filteredPosts, nil
}
