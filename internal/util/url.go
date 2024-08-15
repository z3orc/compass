package util

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func CheckUrl(url string) bool {

	resp, err := http.Get(url)
	if err != nil {
		log.Warn("Failed to check url", "url", url, "error", err)
		return false
	}

	if resp.StatusCode != 200 {
		log.Warn("Failed to check url", "url", url, "status_code", resp.StatusCode)
		return false
	}

	return true
}
