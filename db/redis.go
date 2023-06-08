package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func ConnectToRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address and port
		Password: "",               // Redis password, leave empty if no password is set
		DB:       0,                // Redis database number
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %s", err))
	}
	fmt.Println("Successfully connected to Redis")
	return client
}
