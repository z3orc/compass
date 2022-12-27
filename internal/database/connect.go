package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var Ctx = context.Background()
var Client = Connect()

// Connects & returns a redis client
func Connect() (*redis.Client){
	url, err := redis.ParseURL("redis://localhost:6379")
	if err != nil {
		log.Println(err)
	}

	client := redis.NewClient(url)

	return client
}

// Check the state of a redis client
func Check(client *redis.Client) bool {
	_, err := client.Ping(ctx).Result()

	if(err != nil){
		return false
	} else {
		return true
	}
}