package helper

import (
	"errors"
	"regexp"
)

func CheckEmail(email string) (bool, error) {

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return false, errors.New("Erreur lors de la validation de l'adresse e-mail")
	}

	return match, nil
}

func CheckPassword(password string) (bool, error) {
	// Cette expression exige au moins 8 caractères avec au moins une lettre majuscule,
	// une lettre minuscule, un chiffre et un caractère spécial.
	passwordRegex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`

	match, err := regexp.MatchString(passwordRegex, password)
	if err != nil {
		return false, errors.New("Erreur lors de la validation du mot de passe")
	}

	return match, nil

}

func CheckUserName(username string) (bool, error) {

	// Cette expression exige que le pseudo ait entre 5 et 20 caractères alphanumériques.
	usernameRegex := `^[a-zA-Z0-9]{5,20}$`

	match, err := regexp.MatchString(usernameRegex, username)
	if err != nil {
		return false, errors.New("Erreur lors de la validation du pseudo")
	}

	return match, nil

}

//---------Post------------

func CheckTitle(title string) {

}

//--------Comment--------

func CheckContent(content string) {

}

// func HashPassword(password string) string {
// 	hash := sha256.Sum256([]byte(password))
// 	return fmt.Sprintf("%x", hash)
// }
