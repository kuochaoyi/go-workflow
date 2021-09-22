package dao

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	const config string = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"

	sources := fmt.Sprintf(config,
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DATABASE"),
		os.Getenv("PG_SSLMODE"),
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  sources,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
