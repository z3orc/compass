package data

import (
	"encoding/json"

	"github.com/charmbracelet/log"
	"github.com/z3orc/compass/internal/model"
	"github.com/z3orc/compass/internal/util"
)

type manifest struct {
	Latest   map[string]string
	Versions []versionInfo
}

type versionInfo struct {
	Id   string
	Type string
	Url  string
}

type versionMetadata struct {
	Downloads versionDownloads
	Id        string
}

type versionDownloads struct {
	Server struct {
		Sha1 string
		Size int32
		Url  string
	}
}

type PistonDataSource struct {
	url     string
	flavour model.Flavour
}

func NewPistonDataSource() *PistonDataSource {
	return &PistonDataSource{
		url:     "https://piston-meta.mojang.com/mc/game/version_manifest_v2.json",
		flavour: model.FlavourPiston,
	}
}

func (d *PistonDataSource) GetVersion(id string) (*model.Version, DataError) {
	manifest, err := d.fetchManifest()
	if err != nil {
		log.Error("Unable to retrive manifest", "err", err)
		return nil, &MissingManifestError{}
	}

	var info *versionInfo
	for _, v := range manifest.Versions {
		if v.Id == id {
			info = &v
		}
	}

	if info == nil {
		log.Error("Unable to find version", "id", id, "err", &UnknownVersionError{})
		return nil, &UnknownVersionError{}
	}

	metadata, err := d.fetchMetadata(info)
	if err != nil {
		log.Error("Unable to retrieve metadata about version", "id", id, "err", err)
		return nil, &MissingMetadataError{}
	}

	return &model.Version{
		Flavour: model.FlavourPiston,
		Id:      info.Id,
		Url:     metadata.Downloads.Server.Url,
		Hash:    metadata.Downloads.Server.Sha1,
	}, nil
}

func (d *PistonDataSource) GetFlavour() model.Flavour {
	return d.flavour
}

func (d *PistonDataSource) fetchManifest() (manifest, error) {
	manifest := manifest{}

	res, err := util.GetJson(d.url)
	if err != nil {
		return manifest, err
	}

	err = json.Unmarshal(res, &manifest)
	if err != nil {
		return manifest, err
	}

	return manifest, nil
}

func (d *PistonDataSource) fetchMetadata(info *versionInfo) (versionMetadata, error) {
	version := versionMetadata{}
	res, err := util.GetJson(info.Url)
	if err != nil {
		return version, err
	}

	err = json.Unmarshal(res, &version)
	if err != nil {
		return version, err
	}

	return version, nil
}
