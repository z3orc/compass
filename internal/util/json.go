package util

import (
	"errors"
	"io"
	"net/http"

	"github.com/charmbracelet/log"
)

func GetJson(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Error("Failed to get json", "url", url, "error", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Error("Failed to get json", "url", url, "status_code", resp.StatusCode)
		return nil, errors.New("wrong status code")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return body, err
}

type JsonError struct {
	Error string `json:"message"`
}

func ErrorToJson(err error) JsonError {
	return JsonError{
		Error: err.Error(),
	}
}
