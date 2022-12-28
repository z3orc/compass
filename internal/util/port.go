package util

import "os"

func GetPort() string{
	port := os.Getenv("PORT")

	if port != "" {
		return ":" + port
	}

	return ":8080"
}
