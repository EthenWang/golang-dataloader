package models

type DataLoaderData interface {
	Prepare()
	All() interface{}
	Query(id string) interface{}
	// Update(data DataLoaderDataItem) error
	// Delete(id string) error
}

type Query interface {
	Prepare()
	All() interface{}
	Query(id string) interface{}
}

type ModelCreator interface {
	New() DataLoaderData
}
