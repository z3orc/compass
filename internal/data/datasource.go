package data

import "github.com/z3orc/compass/internal/model"

type IDataSource interface {
	GetVersion(id string) (*model.Version, error)
	GetFlavour() model.Flavour
}
