package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectTiDB() (*gorm.DB, error) {
	dsn := getDSN()

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
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
