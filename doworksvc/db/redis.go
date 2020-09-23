package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func ExampleNewClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rdb.Ping(context.TODO()).Result()
	fmt.Println(pong, err)
}
