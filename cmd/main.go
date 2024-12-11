package main

import (
	"database/sql"
	// "encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Loghadhith/cms/cmd/api"
	"github.com/Loghadhith/cms/db"
	// "github.com/Loghadhith/cms/utils"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {

	c := cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:3000","https://localhost:3000"},        // React frontend
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // HTTP methods
		AllowedHeaders: []string{"Content-Type"},
    AllowCredentials: true,
	})

	db, err := db.NewSQLStorage()

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Hello")

	er := db.Ping()
	if er != nil {
		log.Println("db error fix it")
	}
	log.Println("first db ping")

	initdb(db)

	// r := api.NewAPIServer("localhost:5000", db)
	router := mux.NewRouter()
	api.Run(db, router)
	// router := mux.NewRouter()
	// router.HandleFunc("/",someFunc).Methods("GET")
	// router.Host("http://localhost:5000").Subrouter().HandleFunc("/test",someFunc)
  handler := c.Handler(router)

	srv := http.Server{
		Handler:      handler,
		Addr:         "localhost:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

  srv.ListenAndServe()

	// api.NewAPIServer( "localhost:5000" , db)
	// if err := server.Run(); err != nil {
	//   log.Fatal(err)
	// }
}

// func someFunc(w http.ResponseWriter, r *http.Request) {
//   log.Println("This is default")
//   json.NewEncoder(w).Encode("hello")
//   log.Println(utils.WriteJSON(w,http.StatusOK,nil))
// }

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

	// Execute the CREATE TABLE statement
	_, err = db.Exec(createTableQuery)

	// the pq error is not resolved it is not throwed due to "" or due to '' it is to be further investigated

	if err != nil {
		log.Fatal("Error creating table: ", err)
	}

	log.Println("Table created successfully (if it did not exist already)")
}
