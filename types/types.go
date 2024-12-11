package types

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Pat       string    `json:"pat"`
	CreatedAt time.Time `json:"createdat"`
}

type Post struct {
	ID    int    `json:"id"`
	CID   int    `json:"cid"`
	Repo  string `json:"repo"`
	Fname string `json:"fname"`
	Media string `json:"media"`
}

type PostStore interface {
	PostContentText(txt string) error
	PostMedia() error
}

type UserStore interface {
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(User) error
}

type PostPayload struct {
  Repo string `json:"repo"`
  Path string `json:"file"`
  Type string `json:"ftype"`
  Data string `json:"data"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=5,max=100"`
	Pat      string `json:"pat" validate:"required"`
}
