package main

import (
	"e-commerce/config"
	"e-commerce/db"
	"e-commerce/repository"
	"e-commerce/server"
	"e-commerce/services"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	var cfg = config.GetConfig()

	// connect to db and get the redis instance
	redis := db.ConnectToRedis()

	// connect to db and get the postgres instance
	postgres := db.ConnectToDb()

	// creating repository and injecting postgres instance in it
	repo := repository.NewUserRepository(postgres)

	// creating service and injecting repository in it
	svc := services.NewService(redis, repo)

	r := mux.NewRouter()

	// creating server and injecting service in it
	server.NewServer(svc, cfg, postgres, r)

	_ = fmt.Sprintf("server listening on port %s", cfg.Port)

	err := http.ListenAndServe(cfg.Port, r)
	if err != nil {
		return
	}

}
