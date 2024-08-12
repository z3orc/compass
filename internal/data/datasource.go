package data

type IDataSource interface {
	GetVersion(id string)
}
