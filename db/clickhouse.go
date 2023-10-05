package db

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	gormclickhouse "gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

func ConnectClickhouse() (*gorm.DB, error) {
	sqlDB, err := GetSQLDBLocal()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(gormclickhouse.New(gormclickhouse.Config{
		Conn: sqlDB,
	}))
	if err != nil {
		return nil, err
	}

	// migrate
	err = db.AutoMigrate(&lancrv1.HeroORM{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetSQLDBLocal() (*sql.DB, error) {
	return clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{"localhost:9000"},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 10 * time.Second,
	}), nil
}

func GetSQLDBRemote() (*sql.DB, error) {
	password, ok := os.LookupEnv("CLICKHOUSE_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("CLICKHOUSE_PASSWORD is not set")
	}

	return clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{"z5o0xx7mhd.europe-west4.gcp.clickhouse.cloud:9440"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: password,
		},
		TLS: &tls.Config{
			InsecureSkipVerify: true,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 10 * time.Second,
		Debug:       false,
	}), nil
}
