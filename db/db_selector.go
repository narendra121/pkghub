package db

import (
	postgresdb "github.com/narendra121/pkghub/db/postgres-db"
)

type DbFactory interface {
	Connect() (interface{}, error)
	// CreateTable(tables ...interface{}) error
}

func NewDbFactory(dbf interface{}) DbFactory {
	switch dbf.(type) {
	case *postgresdb.PostgresDB:
		return dbf.(*postgresdb.PostgresDB)
	default:
		return nil
	}
}
