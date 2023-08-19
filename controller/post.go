package controller

import (
	"database/sql"
	"errors"
	"time"

	"forum/models"
)

func CreatePost(db *sql.DB, post models.Post) (int64, error) {
	query := `
        INSERT INTO posts (user_id, title, content, category_id, created_at)
        VALUES (?, ?, ?, ?, ?);
    `

	result, err := db.Exec(query, post.UserID, post.Title, post.Content, post.CategoryID, time.Now())
	if err != nil {
		return 0, err
	}

	postID, _ := result.LastInsertId()
	return postID, nil
}

func GetPostByID(db *sql.DB, postID int64) (models.Post, error) {
	var post models.Post
	query := `
        SELECT id, user_id, title, content, category_id, created_at
        FROM posts
        WHERE id = ?
        LIMIT 1;
    `

	err := db.QueryRow(query, postID).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CategoryID, &post.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Post{}, errors.New("publication non trouvée")
		}
		return models.Post{}, err
	}

	return post, nil
}

func UpdatePost(db *sql.DB, post models.Post) error {
	query := `
        UPDATE posts
        SET title = ?, content = ?, category_id = ?
        WHERE id = ?;
    `

	_, err := db.Exec(query, post.Title, post.Content, post.CategoryID, post.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeletePost(db *sql.DB, postID int64) error {
	query := `
        DELETE FROM posts
        WHERE id = ?;
    `

	_, err := db.Exec(query, postID)
	if err != nil {
		return err
	}

	return nil
}

func GetAllPosts(db *sql.DB) ([]models.Post, error) {
	query := `
        SELECT id, user_id, title, content, category_id, created_at
        FROM posts;
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CategoryID, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
