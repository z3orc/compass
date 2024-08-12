package repo

import "net/http"

type InvalidFlavourError struct{}

func (m *InvalidFlavourError) Error() string {
	return "flavour is not valid/supported"
}

func (m *InvalidFlavourError) StatusCode() uint {
	return http.StatusInternalServerError
}
