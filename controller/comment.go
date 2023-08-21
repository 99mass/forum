package controller

import (
	"database/sql"
	"errors"
	"forum/models"
	"time"

	"github.com/gofrs/uuid"
)

func CreateComment(db *sql.DB, comment models.Comment) (uuid.UUID, error) {
	query := `
        INSERT INTO comments (id, user_id, post_id, content, created_at)
        VALUES (?, ?, ?, ?, ?);
    `

	newUUID, err := uuid.NewV4()
	if err != nil {
		return uuid.UUID{}, err
	}

	_, err = db.Exec(query, newUUID.String(), comment.UserID, comment.PostID, comment.Content, time.Now())
	if err != nil {
		return uuid.UUID{}, err
	}

	return newUUID, nil
}

func GetCommentByID(db *sql.DB, commentID uuid.UUID) (models.Comment, error) {
	var comment models.Comment
	query := `
        SELECT id, user_id, post_id, content, created_at
        FROM comments
        WHERE id = ?
        LIMIT 1;
    `

	err := db.QueryRow(query, commentID).Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content, &comment.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Comment{}, errors.New("commentaire non trouvé")
		}
		return models.Comment{}, err
	}

	return comment, nil
}

func UpdateComment(db *sql.DB, comment models.Comment) error {
	query := `
        UPDATE comments
        SET content = ?
        WHERE id = ?;
    `

	_, err := db.Exec(query, comment.Content, comment.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteComment(db *sql.DB, commentID uuid.UUID) error {
	query := `
        DELETE FROM comments
        WHERE id = ?;
    `

	_, err := db.Exec(query, commentID)
	if err != nil {
		return err
	}

	return nil
}

func GetAllComments(db *sql.DB) ([]models.Comment, error) {
	query := `
        SELECT id, user_id, post_id, content, created_at
        FROM comments;
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

// GetCommentsByPostID retrieves all comments for a specific post by post ID.
func GetCommentsByPostID(db *sql.DB, postID uuid.UUID) ([]models.Comment, error) {
    query := `
        SELECT id, user_id, post_id, content, created_at
        FROM comments
        WHERE post_id = ?;
    `

    rows, err := db.Query(query, postID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var comments []models.Comment
    for rows.Next() {
        var comment models.Comment
        err := rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content, &comment.CreatedAt)
        if err != nil {
            return nil, err
        }
        comments = append(comments, comment)
    }

    return comments, nil
}