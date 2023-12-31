package postgrespkg

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
ppstb := postgresdb.NewPostgresDbBuilder().

		SetUser("narendra").
		SetPassword("123456").
		SetHost("localhost").
		SetDbPort("5432").
		SetDbName("practice").Build()
	dbf := db.NewDbFactory(&ppstb)
	_ = dbf.Connect()
	dbf.CreateTable(&UserPro{}, &RamUsersMegaASD{})
*/
type PostgresFactory interface {
	Connect() (*gorm.DB, error)
}

func NewPostgresFactory(p *PostgresCfg) PostgresFactory {
	return p
}

func (p *PostgresCfg) Connect() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(p.createConnectionString()), &gorm.Config{})

	if err != nil {
		log.Errorln("error in connecting to postgres db :", err)
		return nil, err
	}
	sqlDb, sqlErr := db.DB()
	if sqlErr != nil {
		log.Errorln("error converting postgrs to sql ", sqlErr)
		return nil, sqlErr
	}
	sqlDb.SetMaxIdleConns(5)
	sqlDb.SetMaxOpenConns(20)
	log.Infoln("Postgres is connected", p.createConnectionString())
	return db, nil
}

func (p *PostgresCfg) createConnectionString() string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", p.Host, p.User, p.Password, p.DbName, p.Port)
}
