package database

import (
	"app/src/config"

	"fmt"
	"log"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Connect(dbHost, dbName string) * gorm.DB{
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbHost, config.DBUser, config.DBPassword, dbName, config.DBPort,

	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		fmt.Printf("Failed to connect to database: %+v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from gorm: %+v", err)
	}

	// Config connection pooling
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	return db
}

