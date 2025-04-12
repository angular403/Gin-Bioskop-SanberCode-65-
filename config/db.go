package config

import (
	"github.com/angular403/gin-bioskop/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres:://user:password@localhost:5432/postgres"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
