package main

import (
    "log"
    "os"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func main() {
    dbPath := "./shopping.db"
    dsn := dbPath
    db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database:", err)
    }
    DB = db

    // Migrate
    DB.AutoMigrate(&User{}, &Item{}, &Cart{}, &CartItem{}, &Order{})

    // seed some items if none
    var count int64
    DB.Model(&Item{}).Count(&count)
    if count == 0 {
        DB.Create(&Item{Name: "T-Shirt", Status: "available"})
        DB.Create(&Item{Name: "Mug", Status: "available"})
        DB.Create(&Item{Name: "Notebook", Status: "available"})
    }

    r := SetupRouter()
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}
