package util

import "net/http"

func Error(writer http.ResponseWriter, err error) {
	if err.Error() == "404" {
		http.Error(writer, "404 Not Found -- Could not find requested version", http.StatusNotFound)
	} else if err.Error() == "503" {
		http.Error(writer, "503 Service Unavailable -- Could not reach external resources", http.StatusServiceUnavailable)
	} else {
		http.Error(writer, "400 Bad Request -- Could not process request", http.StatusBadRequest)
	}
	
}