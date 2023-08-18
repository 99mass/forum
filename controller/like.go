package controller

import (
	"database/sql"
	"errors"
	"time"

	"forum/models"

	"github.com/gofrs/uuid"
)

func CreatePostLike(db *sql.DB, like models.PostLike) (uuid.UUID, error) {
	query := `
        INSERT INTO post_likes (id, user_id, post_id, created_at)
        VALUES (?, ?, ?, ?);
    `

	newUUID, err := uuid.NewV4()
	if err != nil {
		return uuid.UUID{}, err
	}

	_, err = db.Exec(query, newUUID.String(), like.UserID, like.PostID, time.Now())
	if err != nil {
		return uuid.UUID{}, err
	}

	return newUUID, nil
}

func CreateCommentLike(db *sql.DB, like models.CommentLike) (uuid.UUID, error) {
	query := `
        INSERT INTO comment_likes (id, user_id, comment_id, created_at)
        VALUES (?, ?, ?, ?);
    `

	newUUID, err := uuid.NewV4()
	if err != nil {
		return uuid.UUID{}, err
	}

	_, err = db.Exec(query, newUUID.String(), like.UserID, like.CommentID, time.Now())
	if err != nil {
		return uuid.UUID{}, err
	}

	return newUUID, nil
}
func GetPostLikeByID(db *sql.DB, likeID uuid.UUID) (models.PostLike, error) {
	var like models.PostLike
	query := `
        SELECT id, user_id, post_id, created_at
        FROM post_likes
        WHERE id = ?
        LIMIT 1;
    `

	err := db.QueryRow(query, likeID).Scan(&like.ID, &like.UserID, &like.PostID, &like.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.PostLike{}, errors.New("like de publication non trouvé")
		}
		return models.PostLike{}, err
	}

	return like, nil
}

func GetCommentLikeByID(db *sql.DB, likeID uuid.UUID) (models.CommentLike, error) {
	var like models.CommentLike
	query := `
        SELECT id, user_id, comment_id, created_at
        FROM comment_likes
        WHERE id = ?
        LIMIT 1;
    `

	err := db.QueryRow(query, likeID).Scan(&like.ID, &like.UserID, &like.CommentID, &like.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.CommentLike{}, errors.New("like de commentaire non trouvé")
		}
		return models.CommentLike{}, err
	}

	return like, nil
}

func UpdatePostLike(db *sql.DB, like models.PostLike) error {
	query := `
        UPDATE post_likes
        SET user_id = ?, post_id = ?
        WHERE id = ?;
    `

	_, err := db.Exec(query, like.UserID, like.PostID, like.ID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCommentLike(db *sql.DB, like models.CommentLike) error {
	query := `
        UPDATE comment_likes
        SET user_id = ?, comment_id = ?
        WHERE id = ?;
    `

	_, err := db.Exec(query, like.UserID, like.CommentID, like.ID)
	if err != nil {
		return err
	}

	return nil
}

func RemovePostLike(db *sql.DB, likeID uuid.UUID) error {
	query := `
        DELETE FROM post_likes
        WHERE id = ?;
    `

	_, err := db.Exec(query, likeID)
	if err != nil {
		return err
	}

	return nil
}

func RemoveCommentLike(db *sql.DB, likeID uuid.UUID) error {
	query := `
        DELETE FROM comment_likes
        WHERE id = ?;
    `

	_, err := db.Exec(query, likeID)
	if err != nil {
		return err
	}

	return nil
}

func GetAllLikes(db *sql.DB) ([]interface{}, error) {
	postLikes, err := GetAllPostLikes(db)
	if err != nil {
		return nil, err
	}

	commentLikes, err := GetAllCommentLikes(db)
	if err != nil {
		return nil, err
	}

	likes := append([]interface{}{}, postLikes)
	likes = append(likes, commentLikes)

	return likes, nil
}

func GetAllPostLikes(db *sql.DB) ([]models.PostLike, error) {
	query := `
        SELECT id, user_id, post_id, created_at
        FROM post_likes;
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var postLikes []models.PostLike
	for rows.Next() {
		var like models.PostLike
		err := rows.Scan(&like.ID, &like.UserID, &like.PostID, &like.CreatedAt)
		if err != nil {
			return nil, err
		}
		postLikes = append(postLikes, like)
	}

	return postLikes, nil
}

func GetAllCommentLikes(db *sql.DB) ([]models.CommentLike, error) {
	query := `
        SELECT id, user_id, comment_id, created_at
        FROM comment_likes;
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commentLikes []models.CommentLike
	for rows.Next() {
		var like models.CommentLike
		err := rows.Scan(&like.ID, &like.UserID, &like.CommentID, &like.CreatedAt)
		if err != nil {
			return nil, err
		}
		commentLikes = append(commentLikes, like)
	}

	return commentLikes, nil
}

func GetPostLikesByPostID(db *sql.DB, postID uuid.UUID) ([]models.PostLike, error) {
    query := `
        SELECT id, user_id, post_id, created_at
        FROM post_likes
        WHERE post_id = ?;
    `

    rows, err := db.Query(query, postID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var likes []models.PostLike
    for rows.Next() {
        var like models.PostLike
        err := rows.Scan(&like.ID, &like.UserID, &like.PostID, &like.CreatedAt)
        if err != nil {
            return nil, err
        }
        likes = append(likes, like)
    }

    return likes, nil
}