package main

import (
	"e-commerce/database"
	"e-commerce/routes"
	"fmt"
	"net/http"
)

func main() {

	database.ConnectToDb()

	r := routes.NewRouter()

	fmt.Println("server listening on port 8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}
