package user

import (
	"database/sql"
	"log"

	"github.com/Loghadhith/cms/types"
	"github.com/Loghadhith/cms/utils"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(user types.User) error {

  _, err := s.db.Exec("INSERT INTO users (name, email, password, pat) VALUES ($1, $2, $3, $4);", 
                         user.Username, user.Email, user.Password, user.Pat)

	if err != nil {
		return err
	}
  log.Println("this is not error so see this is long")

	return nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
  u,r := utils.GetUserByEmail(s.db,email)
  return u,r
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
  u,r := utils.GetUserByID(s.db,id)
  return u,r
}
