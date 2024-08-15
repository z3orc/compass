package model

import "testing"

func TestVersionIsValid(t *testing.T) {
	version := Version{
		Flavour: FlavourPaper,
		Id:      "1.16.5",
		Url:     "https://www.google.com",
		Hash:    "b7a2f3a3f0a8d8c0e0d0c0b0a09080706050403020100",
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
		Hash:    "b7a2f3a3f0a8d8c0e0d0c0b0a09080706050403020100",
	}

	err := version.IsValid()
	if err != nil {
		t.Fatalf("Expected version (%T) to be valid but got error: %v", version, err)
	}

}
