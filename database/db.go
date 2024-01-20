package database

import (
	"log"

	"github.com/felixrodrigo19/rest-api-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DBConection() {
	dsn := "host=localhost user=postgres password=12345 dbname=Novels-API port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("DB connection error")
	}

	DB.AutoMigrate(&models.Author{})
	DB.AutoMigrate(&models.Genre{})
	DB.AutoMigrate(&models.Novel{})
}
