package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectToDb() {
	db, err := sql.Open("postgres", "postgres://davidhorvat@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
}
