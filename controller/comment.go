package controller

import (
	"database/sql"
	"errors"
	"forum/models"
	"time"
)

func CreateComment(db *sql.DB, comment models.Comment) (int64, error) {
	query := `
        INSERT INTO comments (user_id, post_id, content, created_at)
        VALUES (?, ?, ?, ?);
    `

	result, err := db.Exec(query, comment.UserID, comment.PostID, comment.Content, time.Now())
	if err != nil {
		return 0, err
	}

	commentID, _ := result.LastInsertId()
	return commentID, nil
}

func GetCommentByID(db *sql.DB, commentID int64) (models.Comment, error) {
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

func DeleteComment(db *sql.DB, commentID int64) error {
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
