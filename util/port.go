package util

import "os"

func GetPort() string{
	port := os.Getenv("PORT")

	if port == "" {
		return ":8080"
	} else {
		return ":" + port
	}
}
