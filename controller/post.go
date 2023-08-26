package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"forum/models"

	"github.com/gofrs/uuid"
)

func CreatePost(db *sql.DB, post models.Post) (uuid.UUID, error) {

	newUUID, err := uuid.NewV4()
	if err != nil {
		return uuid.UUID{}, err
	}
	fmt.Println("Creating post")
	for _, v := range post.Categories {
		fmt.Println("Creating postCategory", v)
		err := CreatePostCategory(db, newUUID, v.ID)
		if err != nil {
			fmt.Println(err, "pc no cre")
			return v.ID, errors.New("")
		}
	}

	query := `
        INSERT INTO posts (id, user_id, title, content, created_at)
        VALUES (?, ?, ?, ?, ?);
    	`

	_, err = db.Exec(query, newUUID.String(), post.UserID, post.Title, post.Content, time.Now())
	if err != nil {
		fmt.Println(err, "error creating post")
		return uuid.UUID{}, err
	}

	return newUUID, nil
}

func GetPostByID(db *sql.DB, postID uuid.UUID) (models.Post, error) {
	var post models.Post
	query := `
        SELECT id, user_id, title, content, created_at
        FROM posts
        WHERE id = ?
        LIMIT 1;
    `

	err := db.QueryRow(query, postID).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Post{}, errors.New("publication non trouvée")
		}
		return models.Post{}, err
	}
	timeformated := post.CreatedAt[:10]
	post.CreatedAt = timeformated

	return post, nil
}

func UpdatePost(db *sql.DB, post models.Post) error {
	query := `
        UPDATE posts
        SET title = ?, content = ?
        WHERE id = ?;
    `

	_, err := db.Exec(query, post.Title, post.Content, post.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeletePost(db *sql.DB, postID uuid.UUID) error {
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
        SELECT id, user_id, title, content, created_at
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
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		timeformated := post.CreatedAt[:10]
		post.CreatedAt = timeformated
		posts = append(posts, post)
	}

	fmt.Println("time now : ", posts[0].CreatedAt)

	return posts, nil
}

// Function to get user by post ID
func GetUserByPostID(db *sql.DB, postID uuid.UUID) (*models.User, error) {
	query := `
		SELECT u.id, u.username, u.email, u.created_at
		FROM users u
		INNER JOIN posts p ON u.id = p.user_id
		WHERE p.id = ?;
	`

	var user models.User
	err := db.QueryRow(query, postID).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
