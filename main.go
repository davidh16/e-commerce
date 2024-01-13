package main

import (
	"e-commerce/config"
	"e-commerce/db"
	"e-commerce/repository"
	"e-commerce/server"
	"e-commerce/services"
)

func main() {
	var cfg = config.GetConfig()

	// connect to db and get the redis instance
	redis := db.ConnectToRedis()

	// connect to db and get the postgres instance
	postgres := db.ConnectToDb()

	// creating repository and injecting postgres instance in it
	userRepo := repository.NewUserRepository(postgres)
	categoryRepo := repository.NewCategoryRepository(postgres)
	subcategoryRepo := repository.NewSubcategoryRepository(postgres)
	productRepo := repository.NewProductRepository(postgres)
	mediaRepo := repository.NewMediaRepository(postgres)
	addressRepo := repository.NewAddressRepository(postgres)
	paymentInfoRepo := repository.NewPaymentInfoRepository(postgres)
	roleRepo := repository.NewRoleRepository(postgres)
	verificationTokenRepo := repository.NewVerificationTokenRepository(postgres)

	// creating service and injecting repository in it
	svc := services.NewService(
		redis,
		userRepo,
		categoryRepo,
		subcategoryRepo,
		productRepo,
		mediaRepo,
		addressRepo,
		paymentInfoRepo,
		roleRepo,
		verificationTokenRepo,
	)

	// creating server and injecting service in it
	srv := server.NewServer(svc, cfg, redis)

	srv.Serve()

}
