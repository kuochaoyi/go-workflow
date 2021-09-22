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
		os.Getenv("SOURCES_HOST"),
		os.Getenv("SOURCES_PORT"),
		os.Getenv("SOURCES_USER"),
		os.Getenv("SOURCES_PASSWORD"),
		os.Getenv("SOURCES_DATABASE"),
		os.Getenv("SOURCES_SSLMODE"),
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
