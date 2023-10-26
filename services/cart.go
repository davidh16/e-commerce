package services

import (
	"context"
	"e-commerce/models"
	"errors"
	"github.com/mitchellh/mapstructure"
	"strings"
	"time"
)

func (s Service) AddItemToACart(cart *models.Cart) (*models.Cart, error) {
	var updatedCart models.Cart
	cartRedis, err := s.redis.Set(context.Background(), strings.Join([]string{"cart", cart.UserUuid}, "-"), cart, 24*time.Hour).Result()
	if err != nil {
		return nil, errors.New("Failed to store token in Redis")
	}
	err = mapstructure.Decode(cartRedis, &updatedCart) // You can use a library like "mapstructure" for this
	if err != nil {
		return nil, errors.New("Failed to retreive cart from Redis")
	}
	return &updatedCart, nil
}

func (s Service) GetCart(userUuid string) (*models.Cart, error) {
	var cart models.Cart
	cartRedis, err := s.redis.Get(context.Background(), strings.Join([]string{"cart", userUuid}, "-")).Result()
	if err != nil {
		return nil, nil
	}
	err = mapstructure.Decode(cartRedis, &cart) // You can use a library like "mapstructure" for this
	if err != nil {
		return nil, errors.New("Failed to retreive cart from Redis")
	}

	return &cart, nil
}

func (s Service) RemoveItemFromACart(cart *models.Cart) (*models.Cart, error) {
	var updatedCart models.Cart
	cartRedis, err := s.redis.Set(context.Background(), strings.Join([]string{"cart", cart.UserUuid}, "-"), cart, 24*time.Hour).Result()
	if err != nil {
		return nil, errors.New("Failed to store token in Redis")
	}
	err = mapstructure.Decode(cartRedis, &updatedCart) // You can use a library like "mapstructure" for this
	if err != nil {
		return nil, errors.New("Failed to retreive cart from Redis")
	}
	return &updatedCart, nil
}
