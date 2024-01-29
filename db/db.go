package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}
	fmt.Println("Connected to database")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln("Failed to get generic database object from GORM:", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatalln("Failed to close database:", err)
	}
}
