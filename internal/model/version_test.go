package model

import "testing"

func TestVersionIsValid(t *testing.T) {
	version := Version{
		Flavour: FlavourPaper,
		Id:      "1.16.5",
		Url:     "https://www.google.com",
		Hash:    "d8321edc9470e56b8ad5c67bbd16beba25843336",
	}

	err := version.IsValid()
	if err != nil {
		t.Fatalf("Expected version (%T) to be valid but got error: %v", version, err)
	}

}

func TestVersionIsValid2(t *testing.T) {
	version := Version{
		Flavour: FlavourPaper,
		Id:      "1.7",
		Url:     "https://www.google.com",
		Hash:    "90f9c80aeef3966343e661a1487b7918c90ae61d",
	}

	err := version.IsValid()
	if err != nil {
		t.Fatalf("Expected version (%T) to be valid but got error: %v", version, err)
	}

}
