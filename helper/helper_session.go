package helper

import (
	"context"
	"database/sql"
	"forum/controller"
	"forum/models"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

// Example function to create and send a login session cookie
func AddSession(w http.ResponseWriter, userID uuid.UUID, db *sql.DB) {

	expiration := time.Now().Add(24 * time.Hour)
	if userID != uuid.Nil {
		session := models.Session{
			UserID:    userID,
			ExpiresAt: expiration,
			CreatedAt: time.Now(),
		}
		sessionID, err := controller.CreateSession(db, session) // You'll need to implement this function
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Create the session cookie
		cookie := http.Cookie{
			Name:     "sessionID",
			Value:    sessionID.String(),
			Expires:  expiration,
			HttpOnly: true,
			Path:     "/",
		}

		http.SetCookie(w, &cookie)
	}

}

func UpdateSession(db *sql.DB, sessionID uuid.UUID, newExpiration time.Time) error {
	query := `
		UPDATE sessions
		SET expires_at = ?
		WHERE id = ?;
	`

	_, err := db.Exec(query, newExpiration, sessionID)
	if err != nil {
		return err
	}

	return nil
}

func IsEmptySession(s models.Session) bool {
	return s == models.Session{}
}

func GetSessionRequest(r *http.Request) (uuid.UUID, error) {
	// Retrieve the session cookie named "sessionID"
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		// No session cookie foun
		return uuid.Nil, err
	}

	// Extract the value of the session cookie
	sessionid := cookie.Value

	sessionID, err := uuid.FromString(sessionid)
	if err != nil {
		return uuid.Nil, err
	}

	return sessionID, nil

}

func VerifySession(db *sql.DB, sessionID uuid.UUID) bool {
	session, err := controller.GetSessionByID(db, sessionID)
	if err != nil {
		return false
	}
	if &session == nil {
		return false
	}
	return true
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	// Create a new cookie with the same name as the session cookie
	cookie := http.Cookie{
		Name:     "sessionID",
		Value:    "",         // Empty value
		Expires:  time.Now(), // Set to a time in the past
		HttpOnly: true,
		Path:     "/",
	}

	// Set the cookie in the response, effectively deleting it
	http.SetCookie(w, &cookie)
}

// Example function to create and send a login session cookie
func UpdateCookieSession(w http.ResponseWriter, sessionID uuid.UUID, db *sql.DB) {

	expiration := time.Now().Add(24 * time.Hour)

	// Create the session cookie
	cookie := http.Cookie{
		Name:     "sessionID",
		Value:    sessionID.String(),
		Expires:  expiration,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)
	UpdateSession(db, sessionID, expiration)
}

// Middleware to check session and authenticate user
func RequireLogin(next http.Handler, db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the session cookie from the request
		cookie, err := r.Cookie("sessionID")
		if err != nil || cookie == nil {
			// Handle unauthenticated user
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Retrieve the session data from your server-side data store
		sessionIDStr := cookie.Value
		sessionID, err := uuid.FromString(sessionIDStr)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		session, err := controller.GetSessionByID(db, sessionID) // You'll need to implement this function
		if err != nil || IsEmptySession(session) {
			// Handle invalid or expired session
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Attach user ID to the request context for later use
		ctx := context.WithValue(r.Context(), "userID", session.UserID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
