package handler

import (
	"database/sql"
	"forum/controller"
	"forum/helper"
	"forum/models"
	"net/http"
	"strings"
)

func Search(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			title := r.FormValue("title")
			title = strings.TrimSpace(title)
			title = strings.ToLower(title)
			Datas, err := helper.GetDataTemplate(db, r, true, false, false, false, false)
			if err != nil {
				helper.ErrorPage(w, http.StatusBadRequest)
				return
			}
			posts, err := controller.GetPostsByTitle(db, title)
			if err != nil {
				Datas.Error = "there's no post for this title"
			}
			var postsDetails []models.HomeDataPost
			for _, v := range posts {
				post, err := helper.GetDetailPost(db, v)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
					return
				}
				v.Categories,err = controller.GetCategoriesByPost(db,v.ID)
				if err != nil {
					helper.ErrorPage(w,http.StatusBadRequest)
					return
				}
				postsDetails = append(postsDetails, post)
			}
			Datas.Datas = postsDetails
			helper.RenderTemplate(w, "post", "posts", Datas)
		}
	}
}
