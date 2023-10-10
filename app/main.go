package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	response, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(response)

}
