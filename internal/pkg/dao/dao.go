package dao

import (
	"fmt"
	"os"

	config "github.com/kuochaoyi/go-workflow/workflow-config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Model 其它数据结构的公共部分
type Model struct {
	ID int `gorm:"primary_key" json:"id,omitempty"`
}

// 配置
var conf = *config.Config

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
