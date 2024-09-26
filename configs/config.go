package configs

import (
    "github.com/grupoG/csw24-grupoG-ticket-gin/models"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Falha ao conectar ao banco de dados:", err)
    }

    // Migrar o schema
    db.AutoMigrate(&models.Sample{})

    return db
}
