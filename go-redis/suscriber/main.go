package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.99.100:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
	// Publish message
	err = client.Publish(ctx, "mychannel", "payload").Err()
	if err != nil {
		panic(err)
	}
	// Subscribe channel
	pubsub := client.Subscribe(ctx, "mychannel")
	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}
