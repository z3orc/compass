package util

import (
	"errors"
	"io"
	"net/http"
)

func GetJson(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		err = errors.New("503")
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return body, err
}

type JsonError struct {
	Error string `json:"error"`
}

func ErrorToJson(err error) JsonError {
	return JsonError{
		Error: err.Error(),
	}

}
