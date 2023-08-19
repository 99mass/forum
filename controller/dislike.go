package controller

import (
	"database/sql"
	"errors"
	"time"

	"forum/models"
)

func CreatePostDislike(db *sql.DB, dislike models.PostDislike) (int64, error) {
	query := `
        INSERT INTO post_dislikes (user_id, post_id, created_at)
        VALUES (?, ?, ?);
    `

	result, err := db.Exec(query, dislike.UserID, dislike.PostID, time.Now())
	if err != nil {
		return 0, err
	}

	dislikeID, _ := result.LastInsertId()
	return dislikeID, nil
}

func CreateCommentDislike(db *sql.DB, dislike models.CommentDislike) (int64, error) {
	query := `
        INSERT INTO comment_dislikes (user_id, comment_id, created_at)
        VALUES (?, ?, ?);
    `

	result, err := db.Exec(query, dislike.UserID, dislike.CommentID, time.Now())
	if err != nil {
		return 0, err
	}

	dislikeID, _ := result.LastInsertId()
	return dislikeID, nil
}

func GetPostDislikeByID(db *sql.DB, dislikeID int64) (models.PostDislike, error) {
	var dislike models.PostDislike
	query := `
        SELECT id, user_id, post_id, created_at
        FROM post_dislikes
        WHERE id = ?
        LIMIT 1;
    `

	err := db.QueryRow(query, dislikeID).Scan(&dislike.ID, &dislike.UserID, &dislike.PostID, &dislike.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.PostDislike{}, errors.New("dislike de publication non trouvé")
		}
		return models.PostDislike{}, err
	}

	return dislike, nil
}

func GetCommentDislikeByID(db *sql.DB, dislikeID int64) (models.CommentDislike, error) {
	var dislike models.CommentDislike
	query := `
        SELECT id, user_id, comment_id, created_at
        FROM comment_dislikes
        WHERE id = ?
        LIMIT 1;
    `

	err := db.QueryRow(query, dislikeID).Scan(&dislike.ID, &dislike.UserID, &dislike.CommentID, &dislike.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.CommentDislike{}, errors.New("dislike de commentaire non trouvé")
		}
		return models.CommentDislike{}, err
	}

	return dislike, nil
}

func UpdatePostDislike(db *sql.DB, dislike models.PostDislike) error {
	query := `
        UPDATE post_dislikes
        SET user_id = ?, post_id = ?
        WHERE id = ?;
    `

	_, err := db.Exec(query, dislike.UserID, dislike.PostID, dislike.ID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCommentDislike(db *sql.DB, dislike models.CommentDislike) error {
	query := `
        UPDATE comment_dislikes
        SET user_id = ?, comment_id = ?
        WHERE id = ?;
    `

	_, err := db.Exec(query, dislike.UserID, dislike.CommentID, dislike.ID)
	if err != nil {
		return err
	}

	return nil
}

func RemovePostDislike(db *sql.DB, dislikeID int64) error {
	query := `
        DELETE FROM post_dislikes
        WHERE id = ?;
    `

	_, err := db.Exec(query, dislikeID)
	if err != nil {
		return err
	}

	return nil
}

func RemoveCommentDislike(db *sql.DB, dislikeID int64) error {
	query := `
        DELETE FROM comment_dislikes
        WHERE id = ?;
    `

	_, err := db.Exec(query, dislikeID)
	if err != nil {
		return err
	}

	return nil
}

func GetAllDislikes(db *sql.DB) ([]interface{}, error) {
	postDislikes, err := GetAllPostDislikes(db)
	if err != nil {
		return nil, err
	}

	commentDislikes, err := GetAllCommentDislikes(db)
	if err != nil {
		return nil, err
	}

	dislikes := append([]interface{}{}, postDislikes)
	dislikes = append(dislikes, commentDislikes)

	return dislikes, nil
}

func GetAllPostDislikes(db *sql.DB) ([]models.PostDislike, error) {
	query := `
        SELECT id, user_id, post_id, created_at
        FROM post_dislikes;
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var postDislikes []models.PostDislike
	for rows.Next() {
		var dislike models.PostDislike
		err := rows.Scan(&dislike.ID, &dislike.UserID, &dislike.PostID, &dislike.CreatedAt)
		if err != nil {
			return nil, err
		}
		postDislikes = append(postDislikes, dislike)
	}

	return postDislikes, nil
}

func GetAllCommentDislikes(db *sql.DB) ([]models.CommentDislike, error) {
	query := `
        SELECT id, user_id, comment_id, created_at
        FROM comment_dislikes;
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commentDislikes []models.CommentDislike
	for rows.Next() {
		var dislike models.CommentDislike
		err := rows.Scan(&dislike.ID, &dislike.UserID, &dislike.CommentID, &dislike.CreatedAt)
		if err != nil {
			return nil, err
		}
		commentDislikes = append(commentDislikes, dislike)
	}

	return commentDislikes, nil
}
