package db

import (
	"fmt"
	"os"

	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	return ConnectTiDB()
}

func ConnectTiDB() (*gorm.DB, error) {
	dsn := getDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&lancrv1.HeroORM{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDSN() string {
	tidbHost := GetEnvWithDefault("TIDB_HOST", "127.0.0.1")
	tidbPort := GetEnvWithDefault("TIDB_PORT", "4000")
	tidbUser := GetEnvWithDefault("TIDB_USER", "root")
	tidbPassword := GetEnvWithDefault("TIDB_PASSWORD", "")
	tidbDBName := GetEnvWithDefault("TIDB_DB_NAME", "test")
	// TODO implement SSL, notice how currently false and insecure
	useSSL := GetEnvWithDefault("USE_SSL", "false")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&tls=%s",
		tidbUser, tidbPassword, tidbHost, tidbPort, tidbDBName, useSSL)
}

func GetEnvWithDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
