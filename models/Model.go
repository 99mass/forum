package models

import "time"

type User struct {
	ID        int64     `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type Category struct {
	ID           int64  `db:"id"`
	NameCategory string `db:"name_category"`
}

type Post struct {
	ID         int64     `db:"id"`
	UserID     int64     `db:"user_id"`
	Title      string    `db:"title"`
	Content    string    `db:"content"`
	CategoryID int64     `db:"category_id"`
	CreatedAt  time.Time `db:"created_at"`
}

type Comment struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	PostID    int64     `db:"post_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

type PostLike struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	PostID    int64     `db:"post_id"`
	CreatedAt time.Time `db:"created_at"`
}

type PostDislike struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	PostID    int64     `db:"post_id"`
	CreatedAt time.Time `db:"created_at"`
}

type CommentLike struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	CommentID int64     `db:"comment_id"`
	CreatedAt time.Time `db:"created_at"`
}

type CommentDislike struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	CommentID int64     `db:"comment_id"`
	CreatedAt time.Time `db:"created_at"`
}
