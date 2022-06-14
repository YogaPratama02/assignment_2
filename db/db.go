package db

import (
	"BootcampHacktiv8/assignment_2/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbInstance := os.Getenv("DB_INSTANCE")
	dbConnType := os.Getenv("DB_CONNTYPE")

	var err error

	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbConnType, dbInstance, dbname)
	sequel, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("Cannot open sql")
	}

	DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sequel,
	}), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("Cannot connect database")
	}

	sqlDB, _ := DB.DB()
	err = sqlDB.Ping()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Database not responding with err: %s", err)
	}
	DB.AutoMigrate(
		&models.Order{},
		&models.Item{},
	)
}

func DbManager() *gorm.DB {
	return DB
}
