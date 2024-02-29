package purpur

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/z3orc/compass/internal/models"
	"github.com/z3orc/compass/internal/util"
)

const baseURL = "https://api.purpurmc.org/v2/purpur"

type Versions struct {
	Versions []string
}

type Version struct {
	Builds  Builds
	Project string
	Version string
}

type VersionInfo struct {
	Version string
	Md5     string
}

type Builds struct {
	All    []string
	Latest string
}

func GetVersions() (Versions, error) {
	var versions Versions

	resp, err := util.GetJson(baseURL)
	if err != nil {
		return versions, err
	}

	err = json.Unmarshal(resp, &versions)
	if err != nil {
		return versions, err
	}

	return versions, nil
}

func GetVersion(id string) (Version, error) {
	var version Version
	var url string

	versions, err := GetVersions()
	if err != nil {
		return version, err
	}

	length := len(versions.Versions)

	for i := 0; i < int(length); i++ {
		currentId := versions.Versions[i]

		if currentId == id {
			url = fmt.Sprintf("%s/%s", baseURL, currentId)
			break
		}
	}

	if url == "" {
		err := errors.New("404")
		return version, err
	}

	resp, err := util.GetJson(url)
	if err != nil {
		return version, err
	}

	err = json.Unmarshal(resp, &version)
	if err != nil {
		return version, err
	}

	return version, nil
}

func GetLatestBuild(id string) (string, error) {
	version, err := GetVersion(id)
	if err != nil {
		return "", err
	}

	return version.Builds.Latest, nil
}

func GetDownloadUrl(id string) (string, error) {
	latestBuild, err := GetLatestBuild(id)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s/%s/download", baseURL, id, latestBuild)

	return url, nil
}

func GetFormatted(id string) (models.Version, error) {
	latestBuildID, err := GetLatestBuild(id)
	if err != nil {
		return models.Version{}, err
	}

	url, err := GetDownloadUrl(id)
	if err != nil {
		return models.Version{}, err
	}

	latestBuildURL := fmt.Sprintf("%s/%s/%s", baseURL, id, latestBuildID)

	latestBuild, err := util.GetJson(latestBuildURL)
	if err != nil {
		return models.Version{}, err
	}

	var latestBuildJSON = VersionInfo{}
	json.Unmarshal(latestBuild, &latestBuildJSON)

	version := models.Version{
		Url:          url,
		Version:      latestBuildJSON.Version,
		ChecksumType: "md5",
		Checksum:     latestBuildJSON.Md5,
	}

	return version, nil
}
