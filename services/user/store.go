package user

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Loghadhith/cms/types"
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

  stmt := "SELECT * FROM users WHERE"

  if email != "" {
    stmt += fmt.Sprintf(" email = '%v';", email)
  }

	rows, err := s.db.Query(stmt)
	if err != nil {
		return nil, err
	}
  defer rows.Close()

	u := new(types.User)

  defer rows.Close()

	for rows.Next() {

		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1", id)


	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)


	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
    &user.Pat,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
