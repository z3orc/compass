package model

import (
	"strings"
)

type Flavour int

const (
	FlavourPiston Flavour = iota
	FlavourPaper
	FlavourPurpur
)

// ToFlavour converts a string to Flavour (int).
// Returns an int of type Flavour, or -1 if input string is invalid
func StringToFlavour(s string) Flavour {
	s = strings.ToLower(s)

	switch s {
	case "piston":
		return FlavourPiston
	case "paper":
		return FlavourPaper
	case "purpur":
		return FlavourPurpur
	}

	return -1
}

func (f Flavour) ToString() string {
	switch f {
	case FlavourPiston:
		return "piston"
	case FlavourPaper:
		return "paper"
	case FlavourPurpur:
		return "purpur"
	}

	return ""
}

// Checks if the current Flavour is valid
func (f Flavour) IsValid() bool {
	switch f {
	case FlavourPiston, FlavourPaper, FlavourPurpur:
		return true
	}

	return false
}
