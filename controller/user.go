package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"forum/models"

	"github.com/gofrs/uuid"
)

func CreateUser(db *sql.DB, user models.User) (uuid.UUID, error) {
	query := `
        INSERT INTO users (id, username, email, password, created_at)
        VALUES (?, ?, ?, ?, ?);
    `

	newUUID, err := uuid.NewV4()
	if err != nil {
		return uuid.UUID{}, err
	}

	_, err = db.Exec(query, newUUID.String(), user.Username, user.Email, user.Password, time.Now())
	if err != nil {
		return uuid.UUID{}, err
	}

	return newUUID, nil
}

func GetAllUsers(db *sql.DB) ([]models.User, error) {
	query := `
        SELECT id, username, email, password, created_at
        FROM users;
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUserByID retrieves a user by their UUID ID from the database.
func GetUserByID(db *sql.DB, userID uuid.UUID) (*models.User, error) {
	query := `
		SELECT id, username, email, password, created_at
		FROM users
		WHERE id = ?;
	`

	row := db.QueryRow(query, userID)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// No user found for the given ID
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

// GetUserByEmail retrieves a user by their email address from the database.
func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	query := `
		SELECT id, username, email, password, created_at
		FROM users
		WHERE email = ?
		LIMIT 1;
	`

	row := db.QueryRow(query, email)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// No user found for the given email
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func GetUserBySessionId(sessionId uuid.UUID, db *sql.DB) models.User {

	session, _ := GetSessionByID(db, sessionId)
	fmt.Print(session)
	user, _ := GetUserByID(db, session.UserID)
	return *user

}

// GetUserByEmail retrieves a user from the database by email.
// func GetUserByEmail(db *sql.DB, email string) (models.User, error) {
// 	var user models.User
// 	query := `
//         SELECT id, username, email, password, created_at
//         FROM users
//         WHERE email = ?
// 		LIMIT 1;
//     `

// 	stmt, err := db.Prepare(query)
// 	if err != nil {
// 		return user, err
// 	}
// 	defer stmt.Close()

// 	row := stmt.QueryRow(user.Email)
// 	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)

// 	if err != nil {

// 		return user, err
// 	}

// 	return user, nil
// }

func UpdateUser(db *sql.DB, user models.User) error {
	// Mettre à jour uniquement les champs non vides
	query := `
        UPDATE users
        SET username = COALESCE(?, username),
            email = COALESCE(?, email)
        WHERE id = ?;
    `

	_, err := db.Exec(query, user.Username, user.Email, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(db *sql.DB, userID uuid.UUID) error {
	query := `
        DELETE FROM users
        WHERE id = ?;
    `

	_, err := db.Exec(query, userID)
	if err != nil {
		return err
	}

	return nil
}

// Verification duplicatat pseudo ou email.

func IsDuplicateUsernameOrEmail(db *sql.DB, username, email string) (bool, error) {
	query := `
        SELECT COUNT(*)
        FROM users
        WHERE username = ? OR email = ?;
    `

	var count int
	err := db.QueryRow(query, username, email).Scan(&count)
	//fmt.Println(err, ":duplicate",count)
	if err != nil {
		fmt.Println("Error Base de donné")
		return false, errors.New("")
	}
	if count > 0 {
		return true, errors.New("l'utilisateur existe déjà")
	}

	return false, errors.New("")
}

// Function to get user by username
func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	query := `
		SELECT id, username, email, password, created_at
		FROM users
		WHERE username = ?
		LIMIT 1;
	`

	var user models.User
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
