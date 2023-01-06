package env

import (
	"log"
	"os"
)

func ListenerPort() string{
	port := os.Getenv("PORT")

	if port != "" {
		return ":" + port
	}

	return ":8000"
}

func RedisHost() string {
	host := os.Getenv("REDISHOST")

	if host != "" {
		return host
	}

	return "localhost"
}

func RedisPort() string{
	port := os.Getenv("REDISPORT")

	if port != "" {
		return port
	}

	return "6379"
}

func RedisUser() string {
	user := os.Getenv("REDISUSER")

	if user != "" {
		return user
	}

	return "default"
}

func RedisPassword() string {
	password := os.Getenv("REDISPASSWORD")

	if password != "" {
		return password
	}

	return ""
}

func APIURL() string {
    url := os.Getenv("API")

    if url == "" {
        log.Panic("No API url provided, shutting down")
    }

    return url
}
