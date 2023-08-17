package controller

import (
	"database/sql"
	"errors"
	"time"

	"forum/helper"
	"forum/models"
)

func CreateUser(db *sql.DB, user models.User) (int64, error) {
	//Verifier si l'utilisateur exixte deja par email
	existingUser, err := GetUserByEmail(db, user.Email)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if existingUser.ID != 0 {
		return 0, errors.New("l'utilisateur avec cet email existe deja")
	}

	hashedPassword := helper.HashPassword(user.Password)

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

// GetAllUsers retrieves all users from the database.
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

// GetUserByID retrieves a user from the database by ID.
func GetUserByID(db *sql.DB, userID int64) (models.User, error) {
	var user models.User
	query := `
        SELECT id, username, email, password, created_at
        FROM users
        WHERE id = ?
		LIMIT 1;
    `

	err := db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("Utilisateur non trouvé")
		}
		return models.User{}, err
	}

	return user, nil
}

// GetUserByEmail retrieves a user from the database by email.
func GetUserByEmail(db *sql.DB, email string) (models.User, error) {
	var user models.User
	query := `
        SELECT id, username, email, password, created_at
        FROM users
        WHERE email = ?
		LIMIT 1;
    `

	err := db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("Utilisateur non trouvé")
		}
		return models.User{}, err
	}

	return user, nil
}

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

func DeleteUser(db *sql.DB, userID int64) error {
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
