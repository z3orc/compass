package database

import "os"

func GetHost() string {
	host := os.Getenv("REDISHOST")

	if host != "" {
		return host
	}

	return "localhost"
}

func GetPort() string{
	port := os.Getenv("REDISPORT")

	if port != "" {
		return port
	}

	return "6379"
}

func GetUser() string {
	user := os.Getenv("REDISUSER")

	if user != "" {
		return user
	}

	return "default"
}

func GetPassword() string {
	password := os.Getenv("REDISPASSWORD")

	if password != "" {
		return password
	}

	return ""
}