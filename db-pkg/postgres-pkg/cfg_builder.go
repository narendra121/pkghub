package postgrespkg

type PostgresCfg struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
}

type PostgresCfgBuilder struct {
	postgresCfg PostgresCfg
}

func NewPostgresCfgBuilder() *PostgresCfgBuilder {
	return &PostgresCfgBuilder{postgresCfg: PostgresCfg{}}
}

func (db *PostgresCfgBuilder) SetHost(host string) *PostgresCfgBuilder {
	db.postgresCfg.Host = host
	return db
}

func (db *PostgresCfgBuilder) SetUser(userName string) *PostgresCfgBuilder {
	db.postgresCfg.User = userName
	return db
}

func (db *PostgresCfgBuilder) SetPassword(pass string) *PostgresCfgBuilder {
	db.postgresCfg.Password = pass
	return db
}

func (db *PostgresCfgBuilder) SetDbName(dbName string) *PostgresCfgBuilder {
	db.postgresCfg.DbName = dbName
	return db
}

func (db *PostgresCfgBuilder) SetDbPort(dbPort string) *PostgresCfgBuilder {
	db.postgresCfg.Port = dbPort
	return db
}

func (db *PostgresCfgBuilder) Build() PostgresCfg {
	return db.postgresCfg
}
