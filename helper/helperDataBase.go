package helper

import (
	"database/sql"
	"fmt"
	"forum/controller"
	"forum/models"
	"io/ioutil"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database/forum.db")
	if err != nil {
		return nil, err
	}
	// defer db.Close()
	return db, nil
}

func CreateTables(db *sql.DB) error {
	schema, err := ioutil.ReadFile("./database/structure.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}

func Comment(content string, id int) {

}

// The function give all data the the template page needs
func GetDataTemplate(db *sql.DB, r *http.Request, User, Post, Posts, ErrAuth, Category bool) (models.Home, error) {
	var datas models.Home

	//---Get All Posts-------//
	if Posts {
		posts, err := GetPostForHome(db)
		if err != nil {
			fmt.Println("err: ", err)
			//ErrorPage(w, http.StatusInternalServerError)
			return datas, err
		}
		datas.Datas = posts
	}

	//---Get One Post-------//
	if Post {

		ID, err := StringToUuid(r, "post_id")

		if err != nil {
			//ErrorPage(w, http.StatusInternalServerError)
			return datas, err
		}
		postData, errPD := GetPostDetails(db, ID)
		if errPD != nil {
			//ErrorPage(w, http.StatusInternalServerError)
			return datas, err
		}
		datas.PostData = postData
	}

	//---Get the User-------//
	if User {
		var sessiondata bool
		sessionID, errsess := GetSessionRequest(r)
		if errsess != nil {
			sessiondata = false
		} else {
			sessiondata = true
			session, errgets := controller.GetSessionByID(db, sessionID)
			if errgets != nil || &session == nil {
				sessiondata = false
			}
			datas.User = controller.GetUserBySessionId(sessionID, db)
		}
		datas.Session = sessiondata
	}

	//---Get All Categories-------//
	if Category {
		category, err := controller.GetAllCategories(db)
		if err != nil {
			return datas, err
			//ErrorPage(w, http.StatusInternalServerError)
		}
		datas.Category = category
	}

	//---Get Error autification---//
	if ErrAuth {
		email := r.FormValue("email")
		password := r.FormValue("motdepasse")

		okEmail, errE := CheckEmail(email)
		if !okEmail {
			datas.ErrorAuth.EmailError = errE.Error()
			//RenderTemplate(w, "signin", "auth", datas)
			return datas, errE
		} else {
			datas.ErrorAuth.EmailError = ""
		}
		//Check if the error has to be handled
		userID, toConnect := VerifUser(db, email, password)

		if toConnect {
			datas.User.ID = userID
			// Create a session
			//AddSession(w, userID, db)
			// Redirect to home
			//http.Redirect(w, r, "/", http.StatusSeeOther)
			return datas, nil
		} else {
			datas.ErrorAuth.GeneralError = "L'email ou le mot de passe n'est pas correcte"
			return datas, nil
			//RenderTemplate(w, "signin", "auth", datas)
		}
	}

	return datas, nil
}
