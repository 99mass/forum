package controller

import (
	"database/sql"
	"forum/models"
)

// CreatePost inserts a new post into the database.
func CreatePost(db *sql.DB, post models.Post) (int64, error) {
	query := `
        INSERT INTO posts (user_id, title, content, category_id)
        VALUES (?, ?, ?, ?);
    `

	result, err := db.Exec(query, post.UserID, post.Title, post.Content, post.CategoryID)
	if err != nil {
		return 0, err
	}

	postID, _ := result.LastInsertId()
	return postID, nil
}

// GetPostByID retrieves a post from the database by ID.
func GetPostByID(db *sql.DB, postID int64) (models.Post, error) {
	var post models.Post
	query := `
        SELECT id, user_id, title, content, category_id, created_at
        FROM posts
        WHERE id = ?;
    `

	err := db.QueryRow(query, postID).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CategoryID, &post.CreatedAt)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

// GetPostsByCategory retrieves all posts from a specific category.
func GetPostsByCategory(db *sql.DB, categoryID int64) ([]models.Post, error) {
	query := `
        SELECT id, user_id, title, content, category_id, created_at
        FROM posts
        WHERE category_id = ?;
    `

	rows, err := db.Query(query, categoryID)
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

// UpdatePost updates the content of a post in the database.
func UpdatePost(db *sql.DB, postID int64, newContent string) error {
    query := `
        UPDATE posts
        SET content = ?
        WHERE id = ?;
    `

    _, err := db.Exec(query, newContent, postID)
    return err
}

// DeletePost deletes a post from the database.
func DeletePost(db *sql.DB, postID int64) error {
    query := `
        DELETE FROM posts
        WHERE id = ?;
    `

    _, err := db.Exec(query, postID)
    return err
}