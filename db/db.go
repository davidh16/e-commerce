package db

import (
	"e-commerce/config"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var cfg = config.GetConfig()

func ConnectToDb() *gorm.DB {

	db, err := gorm.Open(postgres.Open(cfg.PgUrl), &gorm.Config{})
	if err != nil {
		fmt.Println(cfg.PgUrl)
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Successfully connected to db!")

	return db
}
