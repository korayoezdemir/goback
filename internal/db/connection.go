package db

import (
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once

func GetDB() *gorm.DB {
    once.Do(func() {
        var err error
        dsn := "host=localhost user=user password=password dbname=mydb port=5432 sslmode=disable TimeZone=UTC"
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
            log.Fatalf("Failed to connect to database: %v", err)
        }
    })
    return db
}
