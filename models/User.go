package models

import (
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

func (user *User) CreateUser(db *sql.DB) (int64, error) {
	//Verifier si l'utilisateur exixte deja par email
	existingUser, err := user.GetUserByEmail(db)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if existingUser.ID != 0 {
		return 0, errors.New("l'utilisateur avec cet email existe deja")
	}

	//hashedPassword := helper.HashPassword(user.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	query := `
        INSERT INTO users (username, email, password, created_at)
        VALUES (?, ?, ?, ?);
    `

	result, err := db.Exec(query, user.Username, user.Email, hashedPassword, time.Now())
	if err != nil {
		return 0, err
	}
	userID, _ := result.LastInsertId()
	return userID, nil
}

func (user *User) ConnectUser(db *sql.DB) bool {

	newUser, err := user.GetUserByEmail(db)

	if err != nil {
		return false
	}

	if user.Password == newUser.Password {
		return true
	}

	return false

}

// GetUserByEmail retrieves a user from the database by email.
func (user *User) GetUserByEmail(db *sql.DB) (User, error) {
	var newUser User
	query := `
        SELECT id, username, email, password, created_at
        FROM users
        WHERE email = ?
		LIMIT 1;
    `

	err := db.QueryRow(query, newUser.Email).Scan(&newUser.ID, &newUser.Username, &newUser.Email, &newUser.Password, &newUser.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, errors.New("Utilisateur non trouvé")
		}
		return User{}, err
	}

	return newUser, nil
}

func (user *User) IsDuplicated(db *sql.DB) bool {
	query := `
		SELECT COUNT(*)
		FROM users
		WHERE username = ? OR email = ?;
	`

	var count int
	err := db.QueryRow(query, user.Username, user.Email)
	if err != nil {
		return false
	}

	return count > 0
}

// Register a user
// This function will try to register a user
// It returns true or false weither the register succeeded of not

func (user *User) Register(db *sql.DB) (bool, error) {

	// Check Duplicated case
	if user.IsDuplicated(db) {
		return false, errors.New("Nom d'utilisateur ou adresse e-mail déjà pris")
	}

	_, err := user.CreateUser(db)
	if err != nil {
		return false, errors.New("Erreur lors de l'enregistrement de l'utilisateur")
	}

	return true, nil
}
