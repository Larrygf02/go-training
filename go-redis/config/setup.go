package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewClientRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.99.100:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
	return client
}
