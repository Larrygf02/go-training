package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("Go redis Tutorial")
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.99.100:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
	err = client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	json, err := json.Marshal(Author{Name: "Raul", Age: 26})
	if err != nil {
		fmt.Println(err)
	}
	err = client.Set(ctx, "author", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	author, err := client.Get(ctx, "author").Result()
	fmt.Println("author", author)
}
