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
	// Generate a unique session ID using UUID
	sessionID, err := uuid.NewV4()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set session expiration time (e.g., 1 day from now)
	expiration := time.Now().Add(24 * time.Hour)

	// Create the session cookie
	cookie := http.Cookie{
		Name:     "sessionID",
		Value:    sessionID.String(),
		Expires:  expiration,
		HttpOnly: true, // Prevent JavaScript access
		Path:     "/",  // Cookie is valid for all paths
	}

	// Set the cookie in the response
	http.SetCookie(w, &cookie)

	// Store the session ID and user ID in your server-side data store (e.g., database)
	// Here, you would use your database connection (db) to insert the session into your sessions table
	session := models.Session{
		ID:        sessionID,
		UserID:    userID,
		ExpiresAt: expiration,
		CreatedAt: time.Now(),
	}
	_, err = controller.CreateSession(db, session) // You'll need to implement this function
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
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

func IsEmptySession(s models.Session) bool {
	return s == models.Session{}
}

//requireLoginMiddleware := RequireLogin(yourNextHandler, yourDB)
//router := mux.NewRouter()

// Use the RequireLogin middleware for routes that require authentication
//router.Handle("/protected", requireLoginMiddleware(yourProtectedHandler)).Methods("GET")
