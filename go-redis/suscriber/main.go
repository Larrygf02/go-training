package main

import (
	"context"
	"fmt"

	"github.com/my/repo/config"
)

var ctx = context.Background()

func main() {
	client := config.NewClientRedis()
	// Publish message
	err := client.Publish(ctx, "mychannel", "payload").Err()
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
