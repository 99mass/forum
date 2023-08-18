package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type Session struct {
	ID        string
	UserID    uuid.UUID
	ExpiresAt time.Time
	CreatedAt time.Time
}

type Category struct {
	ID           uuid.UUID `db:"id"`
	NameCategory string    `db:"name_category"`
}

type Post struct {
	ID         uuid.UUID `db:"id"`
	UserID     uuid.UUID `db:"user_id"`
	Title      string    `db:"title"`
	Content    string    `db:"content"`
	CategoryID uuid.UUID `db:"category_id"`
	CreatedAt  time.Time `db:"created_at"`
}

type Comment struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	PostID    uuid.UUID `db:"post_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

type PostLike struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	PostID    uuid.UUID `db:"post_id"`
	CreatedAt time.Time `db:"created_at"`
}

type PostDislike struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	PostID    uuid.UUID `db:"post_id"`
	CreatedAt time.Time `db:"created_at"`
}

type CommentLike struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	CommentID uuid.UUID `db:"comment_id"`
	CreatedAt time.Time `db:"created_at"`
}

type CommentDislike struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	CommentID uuid.UUID `db:"comment_id"`
	CreatedAt time.Time `db:"created_at"`
}
