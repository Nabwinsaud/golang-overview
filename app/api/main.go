package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "nabin"
	dbname   = "godb"
)

var db *sql.DB

func main() {

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	CheckError(err)
	// close database after main function ends
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Successfully connected to database!")

	fmt.Println("Basic app inside !")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
