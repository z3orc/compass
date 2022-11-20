package util

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/z3orc/dynamic-rpc/internal/models"
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

func ReturnJson(w http.ResponseWriter, req *http.Request, version models.Version) {
	jsonBody, _ := json.Marshal(map[string]string{
		"url": version.Url,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBody)
}