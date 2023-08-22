package helper

import (
	"database/sql"
	"fmt"
	"forum/controller"
	"forum/models"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Debug(str string) {
	fmt.Println(str)
}

func GetPostForHome(db *sql.DB) ([]models.HomeDataPost, error) {
	post, err := controller.GetAllPosts(db)
	if err != nil {
		return nil, err
	}
	var HomeDatas []models.HomeDataPost
	for _, post := range post {
		var HomeData models.HomeDataPost
		comments, err := controller.GetCommentsByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		var commentdetails []models.CommentDetails
		for _, com := range comments {
			var commentdetail models.CommentDetails
			commentdetail.Comment = com
			commentlike, err := controller.GetCommentLikesByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			commentdislike, err := controller.GetCommentDislikesByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			commentdetail.CommentLike = len(commentlike)
			commentdetail.CommentDislike = len(commentdislike)
			commentdetails = append(commentdetails, commentdetail)

		}
		likes, err := controller.GetPostLikesByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		nbrlikes := len(likes)
		dislike, err := controller.GetDislikesByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		nbrdislikes := len(dislike)

		HomeData.Posts = post
		HomeData.Comment = commentdetails
		HomeData.PostLike = nbrlikes
		HomeData.PostDislike = nbrdislikes

		HomeDatas = append(HomeDatas, HomeData)
	}
	return HomeDatas, nil
}

func GetPostDetails(db *sql.DB, postID uuid.UUID) (models.HomeDataPost, error) {
	post, err := controller.GetPostByID(db, postID)
	if err != nil {
		return models.HomeDataPost{}, err
	}

	var HomeData models.HomeDataPost
	comments, err := controller.GetCommentsByPostID(db, post.ID)
	if err != nil {
		return models.HomeDataPost{}, err
	}
	var commentdetails []models.CommentDetails
	for _, com := range comments {
		var commentdetail models.CommentDetails
		commentdetail.Comment = com
		commentlike, err := controller.GetCommentLikesByCommentID(db, com.ID)
		if err != nil {
			return models.HomeDataPost{}, err
		}
		commentdislike, err := controller.GetCommentDislikesByCommentID(db, com.ID)
		if err != nil {
			return models.HomeDataPost{}, err
		}
		commentdetail.CommentLike = len(commentlike)
		commentdetail.CommentDislike = len(commentdislike)
		commentdetails = append(commentdetails, commentdetail)

	}
	likes, err := controller.GetPostLikesByPostID(db, post.ID)
	if err != nil {
		return models.HomeDataPost{}, err
	}
	nbrlikes := len(likes)
	dislike, err := controller.GetDislikesByPostID(db, post.ID)
	if err != nil {
		return models.HomeDataPost{}, err
	}
	nbrdislikes := len(dislike)

	HomeData.Posts = post
	HomeData.Comment = commentdetails
	HomeData.PostLike = nbrlikes
	HomeData.PostDislike = nbrdislikes

	return HomeData, nil
}

// func Register(db *sql.DB, user *models.User) (bool, error) {

// 	// Check Duplicated case
// 	dup, err := controller.IsDuplicateUsernameOrEmail(db, user.Username, user.Email)
// 	if dup {
// 		return false, errors.New("Nom d'utilisateur ou adresse e-mail déjà pris")
// 	}

// 	id, err := controller.CreateUser(db, user)
// 	if err != nil {
// 		Debug(err.Error())
// 		return false, errors.New("Erreur lors de l'enregistrement de l'utilisateur")
// 	}
// 	//Debug("Register succeed, Id: " + strconv.FormatInt(id, 10))
// 	return true, nil
// }


func VerifUser(db *sql.DB, email string, password string) (uuid.UUID, bool) {
	user, err := controller.GetUserByEmail(db, email)
	if err != nil {
		return uuid.Nil, false
	}
	if user == nil {
		return uuid.Nil, false
	}
	if !CheckPasswordHash(password, user.Password) {
		return user.ID,false
	}
	return user.ID,true
}

// CheckPasswordHash compares a password with its hashed version
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
