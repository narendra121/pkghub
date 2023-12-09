package db

type DbFactory interface {
	Connect() (interface{}, error)
	// CreateTable(tables ...interface{}) error
}
