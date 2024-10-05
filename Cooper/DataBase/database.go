package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connect() (*gorm.DB, error) {
	dsn := "root:password@tcp(localhost:3306)/cooper_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	return db, nil
}

func Init() error {
	if DB != nil {
		return nil
	}

	db, err := connect()
	if err != nil {
		return err
	}
	DB = db
	return nil
}
