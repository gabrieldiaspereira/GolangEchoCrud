package config

import (
	"os"

	"github.com/gabrieldiaspereira/echoGoApi/domain/product/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if !db.Migrator().HasTable(&models.Product{}) {
		err := db.Migrator().CreateTable(&models.Product{})
		if err != nil {
			panic("failed to create table")
		}
	}
	if err != nil {
		panic("Can't connect to database")
	}
	return db
}
