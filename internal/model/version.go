package model

import (
	"errors"
	"fmt"

	"github.com/z3orc/compass/internal/util"
)

type Version struct {
	Flavour Flavour `json:"flavour"`
	Id      string  `json:"id"`
	Url     string  `json:"url"`
	Hash    string  `json:"hash"`
}

func (v Version) IsValid() error {
	if !v.Flavour.IsValid() {
		return errors.New("invalid flavour")
	}

	if len(v.Id) < 3 || len(v.Id) > 7 {
		fmt.Println(len(v.Id))
		return errors.New("invalid id")
	}

	if !util.CheckUrl(v.Url) {
		return errors.New("invalid url")
	}

	if len(v.Hash) <= 0 {
		return errors.New("invalid hash")
	}

	return nil
}
