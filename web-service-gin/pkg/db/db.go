package db

import (
    "log"

    "github.com/balajisa09/web-service-gin/pkg/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Init() *gorm.DB {
    dbURL := "postgres://pg:pass@localhost:5432/crud"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Album{})

    return db
}