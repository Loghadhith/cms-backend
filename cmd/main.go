package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Loghadhith/cms/cmd/api"
	"github.com/Loghadhith/cms/db"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	db, err := db.NewSQLStorage()

	if err != nil {
		log.Fatal(err)
	}

	er := db.Ping()
	if er != nil {
		log.Println("db error fix it")
	}

	initdb(db)

	router := mux.NewRouter()
	api.Run(db, router)
	handler := c.Handler(router)

	srv := http.Server{
		Handler:      handler,
		Addr:         "localhost:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()

}

func initdb(db *sql.DB) {

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB connected successfully")

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(250) NOT NULL,
    pat VARCHAR(50) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	createType := `DO $$
  BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'content_type') THEN
        CREATE TYPE content_type AS ENUM ('text', 'image');
    END IF;
  END $$;`

	createContentTable := `
  CREATE TABLE IF NOT EXISTS content (
    id SERIAL PRIMARY KEY,
    uid INT NOT NULL,
    repo VARCHAR(255) NOT NULL,
    path varchar(255) NOT NULL,
    type content_type NOT NULL,
    url TEXT NOT NULL,
    createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (uid) REFERENCES users(id)
  );`

	// https://raw.githubusercontent.com/${user_name}/${repo_name}/refs/heads/main/${file_path}

	_, errr := db.Exec(createType)
	_, er := db.Exec(createContentTable)

	_, err = db.Exec(createTableQuery)

	// the pq error is not resolved it is not throwed due to "" or due to '' it is to be further investigated
	//resolved

	if err != nil || er != nil || errr != nil{
		log.Fatal("Error creating table: ", er)
	}

	log.Println("Table created successfully (if it did not exist already)")
}
