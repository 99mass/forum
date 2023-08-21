package models

import (
	"database/sql"
	"errors"
	"forum/helper"
	"strconv"
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

// CreateUser returns the "Id" of the user or "0" and an error
func (user *User) CreateUser(db *sql.DB) (int64, error) {
	//Verifier si l'utilisateur exixte deja par email
	err := user.GetUserByEmail(db)
	if err != sql.ErrNoRows {
		return 0, errors.New("L'utilisateur existe déjà")
	}

	// if existingUser.ID != 0 {
	// 	return 0, errors.New("l'utilisateur avec cet email existe deja")
	// }

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
	clientTOconnect := user
	err := user.GetUserByEmail(db)

	if err != nil {
		return false
	}

	if user.Password == clientTOconnect.Password {
		return true
	}

	return false

}

// GetUserByEmail retrieves a user from the database by email.
// The GetUserByEMail can only be "ErrNoRows" or "nil"
func (user *User) GetUserByEmail(db *sql.DB) error {
	query := `
		SELECT id, username, email, password, created_at
		FROM users
		WHERE email = ?
		LIMIT 1;
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(user.Email)
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		helper.Debug("Utilisateur non trouvé.")
		return err
	}

	return nil
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

	id, err := user.CreateUser(db)
	if err != nil {
		helper.Debug(err.Error())
		return false, errors.New(err.Error())
	}
	helper.Debug("Register succeed, Id: " + strconv.FormatInt(id, 10))
	return true, nil
}
