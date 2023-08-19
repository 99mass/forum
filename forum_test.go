package main

import (
	"fmt"
	"forum/helper"
	"forum/models"
	"testing"
)

var Ch = make(chan error)
var Chh = make(chan error)

var Posts []models.Post
var Comments []models.Comment
var User []models.User
var Category []models.Category


func TestFunc2(t *testing.T) {
	var GetdatCheck = []struct {
		url    string
		model  interface{}
		result bool
	}{
		{
			// "https://groupietrackers.herokuapp.com/api/artists",
			// &Artists,
			// true,
		},
		{
			// "https://groupietrackers.herokuapp.com/api/relation",
			// &Events,
			// true,
		},
		{
			// "https://groupietrackers.herokuapp.com/api/locations",
			// &Location,
			// true,
		},
		{
			// "https://groupietrackers.herokuapp.com/api/date",
			// &Date,
			// false,
		},
	}
	for _, v := range GetdatCheck {
		result := helper.GetJson(v.url, v.model) == nil
		if result == v.result {
			fmt.Println("✅ Test Succeeded ")
		} else {
			t.Error("❌ Test Failed: ")
		}
	}
}
