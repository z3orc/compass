package paper

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/z3orc/dynamic-rpc/internal/models"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

const baseURL = "https://api.papermc.io/v2/projects/paper"

type Versions struct {
	Versions []string
}

type Version struct {
	Builds []int
}

type Build struct {
	Version   string
	Downloads struct {
		Application struct {
			Name   string
			Sha256 string
		}
	}
}

func GetVersions() (Versions, error) {
	resp, err := util.GetJson(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	versions := Versions{}

	err = json.Unmarshal(resp, &versions)
	if err != nil {
		log.Fatal(err)
	}

	return versions, nil
}

func GetVersion(version string) (Version, error) {
	builds := Version{}
    url := fmt.Sprintf("%s/versions/%s", baseURL, version)

	err := util.CheckUrl(url)
	if err != nil {
		return builds, errors.New("404")
	}

	resp, err := util.GetJson(url)
	if err != nil {
		return builds, err
	}

	err = json.Unmarshal(resp, &builds)
	if err != nil {
		return builds, err
	}

	return builds, nil
}

func GetLatestBuild(id string) (string, error) {
	version, err := GetVersion(id)
	if err != nil {
		return "", err
	}
	builds := version.Builds

	latest := builds[len(builds)-1]
	latestAsString := fmt.Sprintf("%v", latest)

	return latestAsString, nil
}


func GetJarName(id string) (string, error) {
	latestBuild, err := GetLatestBuild(id)
	if err != nil {
		return "", err
	}
    url := fmt.Sprintf("%s/versions/%s/builds/%s", baseURL, id, latestBuild)

	resp, err := util.GetJson(url)
	if err != nil {
		return "", err
	}

	build := Build{}

	err = json.Unmarshal(resp, &build)
	if err != nil {
		return "", err
	}

	return build.Downloads.Application.Name, nil
}

func GetDownloadUrl(id string) (string, error) {
	latestBuild, err := GetLatestBuild(id)
	if err != nil {
		return "", err
	}
	jarName, err := GetJarName(id)
	if err != nil {
		return "", err
	}
    url := fmt.Sprintf("%s/versions/%s/builds/%s/downloads/%s", baseURL, id, latestBuild, jarName)
	return url, nil
}

func GetFormatted(id string) (models.Version, error) {
	latestBuild, err := GetLatestBuild(id)
	if err != nil {
		return models.Version{}, err
	}
    url := fmt.Sprintf("%s/versions/%s/builds/%s", baseURL, id, latestBuild)

	resp, err := util.GetJson(url)
	if err != nil {
		return models.Version{}, err
	}

	build := Build{}

	err = json.Unmarshal(resp, &build)
	if err != nil {
		return models.Version{}, err
	}

	url, err = GetDownloadUrl(id)
	if err != nil {
		return models.Version{}, err
	}

    version := models.Version{
        Version: build.Version,
        Url: url,
        ChecksumType: "sha256",
        Checksum: build.Downloads.Application.Sha256,
    }
    
    return version, nil
}
