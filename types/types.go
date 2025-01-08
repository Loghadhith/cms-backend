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
	ID        int       `json:"id"`
	UID       int       `json:"uid"`
	Repo      string    `json:"repo"`
	Path      string    `json:"file"`
	Type      string    `json:"ftype"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type ViewReturn struct {
	File string `json:"path"`
	Type string `json:"type"`
}

type ViewStore interface {
	GetFilesInRepo(repo string, mail string) ([]ViewReturn,error)
}

type PostStore interface {
	PostContent(post PostPayload) error
	PostContentOnExistRepo(post PostPayload) error
	GetPostedData(mail ReqBody) ([]string, error)
}

type UserStore interface {
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(User) error
}

type ReqBody struct {
	Email string `json:"email"`
}

type PostPayload struct {
	Email string `json:"email"`
	Repo  string `json:"repo"`
	Path  string `json:"file"`
	Type  string `json:"ftype"`
	Data  string `json:"data"`
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
