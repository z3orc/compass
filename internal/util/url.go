package util

import (
	"net/http"
)

func CheckUrl(url string) bool {

	resp, err := http.Get(url)
	if err != nil {
		return false
	}

	if resp.StatusCode != 200 {
		return false
	}

	return true
}
