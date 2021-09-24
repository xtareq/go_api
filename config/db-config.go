package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/xtareq/go_api/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DbConnection method for opning db connection
func DbConnection() *gorm.DB {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database connection failded")
	}

	db.AutoMigrate(&entity.User{})

	return db

}

//CloseDb method for closing db Connection
func CloseDb(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic("Something going wrong in database")
	}

	dbSql.Close()
}
