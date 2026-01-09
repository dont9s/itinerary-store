package database

import (
    "log"
    "os"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    db, err := gorm.Open(sqlite.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }
    DB = db
}
