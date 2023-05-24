package database

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/z3orc/dynamic-rpc/internal/env"
)

var RedisCtx = context.Background()

// Connects & returns a redis client
func Connect() *redis.Client {
	const DIAL_TIMEOUT = 100 * time.Millisecond
	const MAX_RETRIES = -1

	var client *redis.Client = nil

	opt, err := redis.ParseURL(fmt.Sprint(env.RedisURL(), fmt.Sprintf("?dial_timeout=%s&max_retries=%d", DIAL_TIMEOUT, MAX_RETRIES)))

	if err == nil {
		client = redis.NewClient(opt)
	} else {
		client = redis.NewClient(&redis.Options{
			Addr:        fmt.Sprint(env.RedisHost(), ":", env.RedisPort()),
			Username:    env.RedisUser(),
			Password:    env.RedisPassword(), // no password set
			DB:          0,                   // use default DB
			DialTimeout: DIAL_TIMEOUT,
			MaxRetries:  MAX_RETRIES,
		})
	}

	return client
}

// // Check the state of a redis client
// func Check(client *redis.Client) bool {
// 	_, err := client.Ping(RedisCtx).Result()

// 	if(err != nil){
// 		return false
// 	} else {
// 		return true
// 	}
// }
