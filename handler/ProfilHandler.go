package handler

import (
	"database/sql"
	"fmt"
	"forum/helper"
	"forum/models"
	"net/http"
)

func GetProfil(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var dataProfil models.DataMyProfil
		catMap := map[string]int{}
		datas, err := helper.GetDataTemplate(db, r, true, false, false, false, true)

		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
			return
		}
		if !datas.Session {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

		PostsDetails, err := helper.GetPostsForOneUser(db, datas.User.ID)
		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
			return
		}

		dataProfil.Posts = PostsDetails
		dataProfil.User = datas.User
		for _, cat := range datas.Category {
			catMap[cat.NameCategory] = 0
		}
		
		for _, post := range dataProfil.Posts {
			for _, cat := range post.Posts.Categories {

				catMap[cat.NameCategory] += 1
				
			}
		}
		dataProfil.Categories = catMap
		datas.DataProfil = dataProfil

		helper.RenderTemplate(w, "profil", "profil", datas)
	}
}
