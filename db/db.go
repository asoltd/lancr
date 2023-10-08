package db

import (
	"os"

	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := ConnectTiDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&lancrv1.HeroORM{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&lancrv1.CourseORM{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&lancrv1.ApprenticeORM{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetEnvWithDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
