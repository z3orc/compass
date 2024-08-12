package data

type MissingManifestError struct{}

func (m *MissingManifestError) Error() string {
	return "unable to retrive manifest"
}

type MissingMetadataError struct{}

func (m *MissingMetadataError) Error() string {
	return "unable to retrive metadata/info about requested version"
}

type UnknownVersionError struct{}

func (m *UnknownVersionError) Error() string {
	return "no version found with given id"
}
