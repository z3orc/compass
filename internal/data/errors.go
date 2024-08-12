package data

import "net/http"

type DataError interface {
	Error() string
	StatusCode() uint
}

type MissingManifestError struct{}

func (m *MissingManifestError) Error() string {
	return "Unable to retrive manifest"
}

func (m *MissingManifestError) StatusCode() uint {
	return http.StatusServiceUnavailable
}

type MissingMetadataError struct{}

func (m *MissingMetadataError) Error() string {
	return "Unable to retrive metadata/info about requested version"
}

func (m *MissingMetadataError) StatusCode() uint {
	return http.StatusServiceUnavailable
}

type UnknownVersionError struct{}

func (m *UnknownVersionError) Error() string {
	return "No version found with given id"
}

func (m *UnknownVersionError) StatusCode() uint {
	return http.StatusNotFound
}
