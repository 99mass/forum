package helper

import (
	"crypto/sha256"
	"fmt"
)

func CheckEmail(email string) {

}

func CheckPassword(password string) {

}

func CheckDuplicateEmail(email string) {

}

func CheckUserName(user string) {

}

//---------Post------------

func CheckTitle(title string) {

}

//--------Comment--------

func CheckContent(content string) {

}

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}
