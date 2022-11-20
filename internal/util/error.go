package util

import "net/http"

func Error(writer http.ResponseWriter, err error) {
	if err.Error() == "404" {
		http.Error(writer, "Could not find version", http.StatusNotFound)
	} else if err.Error() == "503" {
		http.Error(writer, "Service Unavailable", http.StatusServiceUnavailable)
	} else {
		http.Error(writer, "Bad Request", http.StatusBadRequest)
	}
}