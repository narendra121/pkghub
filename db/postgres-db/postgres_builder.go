package postgresdb

type PostgresDB struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
}

type PostgresDbBuilder struct {
	dm PostgresDB
}

func NewPostgresDbBuilder() *PostgresDbBuilder {
	return &PostgresDbBuilder{dm: PostgresDB{}}
}

func (db *PostgresDbBuilder) SetHost(host string) *PostgresDbBuilder {
	db.dm.Host = host
	return db
}

func (db *PostgresDbBuilder) SetUser(userName string) *PostgresDbBuilder {
	db.dm.User = userName
	return db
}

func (db *PostgresDbBuilder) SetPassword(pass string) *PostgresDbBuilder {
	db.dm.Password = pass
	return db
}

func (db *PostgresDbBuilder) SetDbName(dbName string) *PostgresDbBuilder {
	db.dm.DbName = dbName
	return db
}

func (db *PostgresDbBuilder) SetDbPort(dbPort string) *PostgresDbBuilder {
	db.dm.Port = dbPort
	return db
}

func (db *PostgresDbBuilder) Build() PostgresDB {
	return db.dm
}
