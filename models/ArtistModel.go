package models

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id       int
	Name     string
	UserName string
	Country  []string
	Age      int
}
type Post struct {
	Id       int
	Category string
	Title    string
	Date     string
	Flag     string
	Likes    int
	Comments map[string]Comment
}

type Category struct {
	Id     int
	Name   string
	Childs map[string]Category
	Parent string
}

type Comment struct {
	Id      int
	User    int
	Likes   int
	Content string
}

//****************************** GET ARTIST METHODE ******************************************************
// func (ArtistOne) GetAllartists() (bool, []ArtistOne) {

// }

func GetJson(url string, model interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(model)
}
