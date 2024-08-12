package repo

import "github.com/z3orc/compass/internal/model"

type IVersionRepository interface {
	GetVersion(flavour model.Flavour, id string) (*model.Version, error)
	GetFlavours() []model.Flavour
}
