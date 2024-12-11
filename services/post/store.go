package post

import "database/sql"


type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}


func (s *Store) PostContentText(txt string) error {
    // Implement the logic to handle text data
    // For example, inserting it into a database:
    // INSERT INTO posts (text) VALUES (?) 
    return nil
}

// PostMedia should handle media data related to the post
func (s *Store) PostMedia() error {
    // Implement the logic to handle media data
    return nil
}
