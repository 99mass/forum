package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ExpiresAt time.Time
	CreatedAt time.Time
}

type Category struct {
	ID           uuid.UUID
	NameCategory string
}

type Post struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	Title      string
	Content    string
	CategoryID []uuid.UUID
	CreatedAt  string
	//CreatedAt  time.Time
}

type Comment struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	PostID    uuid.UUID
	Content   string
	CreatedAt time.Time
}

type PostLike struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	PostID    uuid.UUID
	CreatedAt time.Time
}

type PostDislike struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	PostID    uuid.UUID
	CreatedAt time.Time
}

type CommentLike struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	CommentID uuid.UUID
	CreatedAt time.Time
}

type CommentDislike struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	CommentID uuid.UUID
	CreatedAt time.Time
}

type HomeDataPost struct {
	Posts       Post
	Comment     []CommentDetails
	PostLike    int
	PostDislike int
	User        User
}

type CommentDetails struct {
	Comment        Comment
	CommentLike    int
	CommentDislike int
	User           User
}

type Home struct {
	Session   bool
	Category  []Category
	Datas     []HomeDataPost
	User      User
	ErrorAuth ErrorAuth
	PostData  HomeDataPost
}

type PostCategory struct {
	CategoryID uuid.UUID
	PostID     uuid.UUID
}

type ErrorAuth struct {
	EmailError    string
	UserNameError string
	PasswordError string
	GeneralError  string
}
