package database

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisCtx = context.Background()

// Connects & returns a redis client
func Connect() (*redis.Client){
	// url, err := redis.ParseURL("redis://localhost:6379")
	// if err != nil {
	// 	log.Println(err)
	// }

	client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
		DialTimeout: 1*time.Second,
		MaxRetries: -1,
    })

	return client
}

// Check the state of a redis client
func Check(client *redis.Client) bool {
	_, err := client.Ping(RedisCtx).Result()

	if(err != nil){
		return false
	} else {
		return true
	}
}