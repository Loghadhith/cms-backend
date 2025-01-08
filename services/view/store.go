package view

import (
	"database/sql"

	"github.com/Loghadhith/cms/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetFilesInRepo(repo string, mail string) ([]types.ViewReturn, error) {
	asjdlkj
}
