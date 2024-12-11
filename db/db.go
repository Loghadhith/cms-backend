package db

import (
  "database/sql"
  "log"
  "fmt"
)


func NewSQLStorage()(*sql.DB , error){

  const (
    host = "localhost"
    port = 5432
    user = "postgres"
    password = "password"
    dbname = "cms"
  )

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    log.Fatal(err)
  }

  return db,nil

}
