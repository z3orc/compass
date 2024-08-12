package data

type MissingManifestError struct{}

func (m *MissingManifestError) Error() string {
	return "unable to retrive manifest"
}

type UnknownVersionIDError struct{}

func (m *UnknownVersionIDError) Error() string {
	return "no version found with given id"
}
