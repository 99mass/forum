package controller

import (
	"database/sql"
	"time"

	"forum/models"
)

// CreateSession creates a new session in the database.
func CreateSession(db *sql.DB, session models.Session) error {
	query := `
        INSERT INTO sessions (id, user_id, expires_at)
        VALUES (?, ?, ?);
    `

	_, err := db.Exec(query, session.ID, session.UserID, session.ExpiresAt)
	return err
}

// GetSessionByID retrieves a session by ID from the database.
func GetSessionByID(db *sql.DB, sessionID string) (models.Session, error) {
	var session models.Session
	query := `
        SELECT id, user_id, expires_at
        FROM sessions
        WHERE id = ?;
    `

	err := db.QueryRow(query, sessionID).Scan(&session.ID, &session.UserID, &session.ExpiresAt)
	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}

// GetSessionUserID retrieves the user ID associated with a session from the database.
func GetSessionUserID(db *sql.DB, sessionID string) (int64, error) {
	query := `
        SELECT user_id
        FROM sessions
        WHERE id = ? AND expires_at > ?;
    `

	var userID int64
	err := db.QueryRow(query, sessionID, time.Now()).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// DeleteSession deletes a session by ID from the database.
func DeleteSession(db *sql.DB, sessionID string) error {
	query := `
        DELETE FROM sessions
        WHERE id = ?;
    `

	_, err := db.Exec(query, sessionID)
	return err
}

// ValidateSession checks if a session is valid based on its expiration time.
func ValidateSession(session models.Session) bool {
	return session.ExpiresAt.After(time.Now())
}

// func CreateSession(userID int64, duration time.Duration) (string, error) {
// 	sessionID, err := GenerateRandomID(32) // Génère un ID de session de 32 octets
//     if err != nil {
//         fmt.Println("Erreur lors de la génération de l'ID de session :", err)
//         return
//     }

// 	expiresAt := time.Now().Add(duration)
// 	createdAt := time.Now()

// 	_, err := db.Exec("INSERT INTO sessions (id, user_id, expires_at, created_at) VALUES (?, ?, ?, ?)",
// 		sessionID, userID, expiresAt, createdAt)
// 	if err != nil {
// 		return "", err
// 	}

// 	return sessionID, nil
// }

// func GenerateRandomID(length int) (string, error) {
// 	bytes := make([]byte, length)
// 	_, err := rand.Read(bytes)
// 	if err != nil {
// 		return "", err
// 	}

// 	return hex.EncodeToString(bytes), nil
// }
