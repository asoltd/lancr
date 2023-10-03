package db

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return ConnectClickhouse()
}

// TODO using local clickhouse for now
// add support for cloud requires running ./clickhouse server, otherwise this
// function will fail
func ConnectClickhouse() (*gorm.DB, error) {
	dsn := "tcp://localhost:9000"
	dialector := clickhouse.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
