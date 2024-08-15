package model

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/charmbracelet/log"
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

	//Minecraft version IDs have been altered throughout the history of Minecraft. Probably not a good idea to check for exact length.
	// if len(v.Id) < 3 || len(v.Id) > 7 {
	// 	fmt.Println(len(v.Id))
	// 	return errors.New("invalid id")
	// }

	if !util.CheckUrl(v.Url) {
		return errors.New("invalid url")
	}

	//FIXME: SHA1 hash is 40 characters long. Need to be changed for other algorithms.
	if len(v.Hash) != 40 {
		return errors.New("invalid hash. supposed to be 40 characters long but was " + strconv.Itoa(len(v.Hash)))
	}

	return nil
}

func (v Version) ToJson() ([]byte, error) {
	res, err := json.Marshal(
		struct {
			Flavour string `json:"flavour"`
			Id      string `json:"id"`
			Url     string `json:"url"`
			Hash    string `json:"hash"`
		}{
			Flavour: v.Flavour.ToString(),
			Id:      v.Id,
			Url:     v.Url,
			Hash:    v.Hash,
		},
	)
	if err != nil {
		log.Error("Failed to marshal version to json", "version", v, "error", err)
		return nil, err
	}

	return res, nil
}
