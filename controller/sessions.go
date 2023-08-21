package controller

import (
	"database/sql"
	"time"

	"forum/models"

	"github.com/gofrs/uuid"
)

func CreateSession(db *sql.DB, session models.Session) (error) {
	query := `
        INSERT INTO sessions (id, user_id, expires_at)
        VALUES (?, ?, ?);
    `

	// newUUID, err := uuid.NewV4()
	// if err != nil {
	// 	return uuid.UUID{}, err
	// }

	_, err := db.Exec(query, session.ID, session.UserID, session.ExpiresAt)
	if err != nil {
		return err
	}

	return  nil
}

// GetSessionByID retrieves a session by ID from the database.
func GetSessionByID(db *sql.DB, sessionID uuid.UUID) (models.Session, error) {
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
func GetSessionUserID(db *sql.DB, sessionID uuid.UUID) (uuid.UUID, error) {
	query := `
        SELECT user_id
        FROM sessions
        WHERE id = ? AND expires_at > ?;
    `

	var userID uuid.UUID
	err := db.QueryRow(query, sessionID, time.Now()).Scan(&userID)
	if err != nil {
		return uuid.UUID{}, err
	}

	return userID, nil
}

// DeleteSession deletes a session by ID from the database.
func DeleteSession(db *sql.DB, sessionID uuid.UUID) error {
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

// func CreateSession(userID uuid.UUID, duration time.Duration) (uuid.UUID, error) {
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

// func GenerateRandomID(length int) (uuid.UUID, error) {
// 	bytes := make([]byte, length)
// 	_, err := rand.Read(bytes)
// 	if err != nil {
// 		return "", err
// 	}

// 	return hex.EncodeToString(bytes), nil
// }
