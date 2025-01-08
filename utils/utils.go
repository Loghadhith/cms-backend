package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Loghadhith/cms/types"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

var Validate = validator.New()

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func ParseJSON(r *http.Request, v any) error {

	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}
	return ""
}

func GetUserByEmail(db *sql.DB, email string) (*types.User, error) {

	stmt := "SELECT * FROM users WHERE"

	if email != "" {
		stmt += fmt.Sprintf(" email = '%v';", email)
	}

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	u := new(types.User)

	defer rows.Close()

	for rows.Next() {

		u, err = ScanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func GetUserByID(db *sql.DB, id int) (*types.User, error) {
	rows, err := db.Query("SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = ScanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func GetPostedData(db *sql.DB, mail types.ReqBody) ([]string, error) {

  log.Println("Enetred Utils fucntion")

	user, err := GetUserByEmail(db, mail.Email)

  id := user.ID

	if err != nil {
		return nil, err
	}

	stmt := "SELECT repo FROM content WHERE uid = "

	stmt += fmt.Sprintf("%v;", id)

	log.Println(stmt)
	log.Println(user.Email)

	rows, err := db.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Use a map to ensure uniqueness of repo names
	repoNamesMap := make(map[string]struct{})

	for rows.Next() {
		var repoName string
		err := rows.Scan(&repoName)
		if err != nil {
			return nil, err
		}
		// Add to the map to ensure uniqueness
		repoNamesMap[repoName] = struct{}{}
	}

	// Check if there was an error in iterating the rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Convert the map to a slice
	var repoNames []string
	for repoName := range repoNamesMap {
		repoNames = append(repoNames, repoName)
	}

  log.Println(repoNames)
  log.Println("return")

	return repoNames, nil

}

func ScanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
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
