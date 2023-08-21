package helper

import (
	"errors"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func CheckRegisterFormat(username, email, password string) error {
	okUserName, errUN := CheckUserName(username)
	okEmail, errE := CheckEmail(email)
	okPassWord, errP := CheckPassword(password)
	if !okUserName {
		//Debug(errUN.Error())
		fmt.Println(errUN)
		return errUN
	}
	if !okEmail {
		Debug(errE.Error())
		return errE
	}
	if !okPassWord {
		Debug(errP.Error())
		return errP
	}
	return nil
}

func CheckEmail(email string) (bool, error) {

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, _ := regexp.MatchString(emailRegex, email)
	if !match {
		return false, errors.New("Format email non valide!")
	}

	return match, nil
}

func CheckPassword(password string) (bool, error) {
	// Cette expression exige au moins 8 caractères avec au moins une lettre majuscule,
	// une lettre minuscule, un chiffre et un caractère spécial.
	// Vérification de la longueur

	if len(password) < 8 || len(password) > 25 {
		fmt.Println("Mot de passe invalide en termes de longueur")
		return false, errors.New("Longueur mot de passe non valide: minimum 8, maximum 25")
	}

	// Vérification des autres conditions avec des expressions régulières
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	digitRegex := regexp.MustCompile(`\d`)
	specialCharRegex := regexp.MustCompile(`[@$!%*?&_\-]`)

	if !lowercaseRegex.MatchString(password) {
		fmt.Println("Le mot de passe doit contenir au moins une lettre minuscule")
		return false, errors.New("Le mot de passe doit contenir au moins une lettre minuscule")
	}
	if !uppercaseRegex.MatchString(password) {
		fmt.Println("Le mot de passe doit contenir au moins une lettre majuscule")
		return false, errors.New("Le mot de passe doit contenir au moins une lettre majuscule")
	}
	if !digitRegex.MatchString(password) {
		fmt.Println("Le mot de passe doit contenir au moins un chiffre")
		return false, errors.New("Le mot de passe doit contenir au moins un chiffre")
	}
	if !specialCharRegex.MatchString(password) {
		fmt.Println("Le mot de passe doit contenir au moins un caractère spécial")
		return false, errors.New("Le mot de passe doit contenir au moins un caractère spécial")
	}

	// if !match {
	// 	return false, errors.New("Format mot de passe non valide!")
	// }

	return true, nil

}

func CheckUserName(username string) (bool, error) {
	Debug("checkusername:" + username)
	// Cette expression exige que le pseudo ait entre 5 et 20 caractères alphanumériques.
	usernameRegex := `^[a-zA-Z0-9]{5,20}$`

	match, _ := regexp.MatchString(usernameRegex, username)

	if !match {
		return false, errors.New("Format du username non valide!")
	}

	return match, nil

}

//---------Post------------

func CheckTitle(title string) {

}

//--------Comment--------

func CheckContent(content string) {

}

// Cryptage du mot de passe
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		Debug("Erreur du hashage")
		return "", err
	}
	Debug("hashage réussi")
	return string(hashedPassword), nil
}

// Confirmation du mot de passe lors de l'inscription d'uun nouveau client
func ConfirmPasswordsMatch(password, confirmPassword string) bool {
	return password == confirmPassword
}
