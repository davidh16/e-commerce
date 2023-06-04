package main

import (
	"e-commerce/config"
	"e-commerce/controller"
	"e-commerce/db"
	"e-commerce/repository"
	"e-commerce/routes"
	"e-commerce/services"
	"fmt"
	"github.com/slayer/autorestart"
	"net/http"
)

func main() {
	autorestart.StartWatcher()

	var cfg = config.GetConfig()

	// connect to db and get the postgres instance
	postgres := db.ConnectToDb()

	// creating repository and injecting postgres instance in it
	repo := repository.NewRepository(postgres)

	// creating service and injecting repository in it
	svc := services.NewService(repo)

	// creating controller and injecting service in it
	ctrl := controller.NewController(svc)

	_ = fmt.Sprintf("server listening on port %s", cfg.Port)

	// creating router and injecting controller in it so all routes have access to handling functions of controller
	r := routes.NewRouter(ctrl)

	err := http.ListenAndServe(cfg.Port, r)
	if err != nil {
		return
	}

}
