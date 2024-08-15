package util

import (
	"errors"
	"testing"
)

func TestGetJson(t *testing.T) {
	_, err := GetJson("https://www.google.com")
	if err != nil {
		t.Error(err)
	}
}

func TestGetJsonError(t *testing.T) {
	_, err := GetJson("https://www.google.com/404")
	if err == nil {
		t.Error("error expected")
	}
}

func TestErrorToJson(t *testing.T) {
	err := errors.New("test")
	json := ErrorToJson(err)
	if json.Error != err.Error() {
		t.Error("error not equal")
	}
}

func TestErrorToJsonString(t *testing.T) {
	json := ErrorToJson(errors.New("test"))
	if json.Error != "test" {
		t.Error("error not equal")
	}
}
