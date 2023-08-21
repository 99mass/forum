package helper

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"forum/controller"
	"forum/models"
	"net/http"
	"strconv"
	"strings"
)

// ******************* VERIF IF THE STRING IS AN INT*****************************************************
func IsInt(s string) bool {
	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

// ***************************** GET JSON DATA ************************************************************
func GetJson(url string, model interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(model)
}

// ******************************* PARSE FILE IN URL *****************
func PArseUlr(r *http.Request, match string) (bool, int) {
	index := strings.Split(r.URL.Path[1:], "/")
	if len(index) == 2 && index[0] == match {
		id, err := strconv.Atoi(index[1])
		if err == nil {
			return true, id
		}
	}
	return false, 0
}

// *************************************** FILTER THE DATA WITH THE PARAM****************
func FilterData() {

}

func FilterCategpry() {

}

func FilterCreationDate() {

}

func FilterLocation() {

}

// *************************************** FECTH ERROR **********************************
func FecthError(ch <-chan error) bool {
	for err := range ch {
		if err != nil {
			fmt.Println(err)
			return true
		}
	}
	return false
}

// ************************************** REMOVE THE DUP *********************************
func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func FilterLoc(_items []string, s string) []string {
	items := []string{}
	for _, v := range _items {
		if strings.Contains(strings.ToLower(v), strings.ToLower(s)) {
			items = append(items, v)
		}
	}
	return items
}

func CheckErr(errs ...error) bool {
	for _, v := range errs {
		if v != nil {
			return true
		}
	}
	return false
}

func GetYear(date string) (int, error) {
	stringyear := strings.Split(date, "-")[2]
	if len(stringyear) == 4 {
		year, err := strconv.Atoi(stringyear)
		return year, err
	}
	return 0, errors.New("len error")
}

func GetPostForHome(db *sql.DB) ([]models.HomeData, error) {
	posts, err := controller.GetAllPosts(db)
	if err != nil {
		return nil, err
	}
	var HomeDatas []models.HomeData
	for _, post := range posts {
		var HomeData models.HomeData
		likes, err := controller.GetPostLikesByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		nbrlikes := len(likes)
		comments, err := controller.GetCommentsByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		var commentdetails []models.CommentDetails
		for _,w := range comments{
			var commentdetail models.CommentDetails
			commentdetail.Comment = w
			commentlike,err := controller.GetCommentLikesByCommentID(db,w.ID)
			if err != nil {
				return nil,err
			}
			commentdislike,err := controller.GetCommentDislikesByCommentID(db,w.ID)
			if err != nil {
				return nil,err
			}
			commentdetail.CommentLike = len(commentlike)
			commentdetail.CommentDislike = len(commentdislike)
			commentdetails = append(commentdetails, commentdetail)

		}
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
