package main

import (
	"e-commerce/config"
	"e-commerce/db"
	"e-commerce/handlers"
	"e-commerce/repository"
	"e-commerce/routes"
	"e-commerce/services"
	"fmt"
	"net/http"
)

func main() {

	var cfg = config.GetConfig()

	// connect to db and get the postgres instance
	postgres := db.ConnectToDb()

	// creating repository and passing postgres instance
	repo := repository.NewRepository(postgres)

	// creating service and passing repository to it
	svc := services.NewService(repo)

	_ = handlers.NewUserHandler(svc)

	r := routes.NewRouter()

	_ = fmt.Sprintf("server listening on port %s", cfg.Port)

	err := http.ListenAndServe(cfg.Port, r)
	if err != nil {
		return
	}

}
