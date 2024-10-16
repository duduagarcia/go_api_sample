package configs

import (
	"fmt"
	"log"

	"github.com/grupoG/csw24-grupoG-ticket-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
    // db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    // if err != nil {
    //     log.Fatal("Falha ao conectar ao banco de dados:", err)
    // }

    // // Migrar o schema
    // db.AutoMigrate(&models.Sample{})

    // return db

    // Obter informações de conexão do ambiente
    // host := os.Getenv("DB_HOST")
    // user := os.Getenv("POSTGRES_USER")
    // password := os.Getenv("POSTGRES_PASSWORD")
    // dbname := os.Getenv("POSTGRES_DB")

    host := "postgresdb"
    user := "postgres"
    password := "postgres"
    dbname := "mydb"
    port := "5432" // Porta padrão do Postgres

    // String de conexão
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Falha ao conectar ao banco de dados:", err)
    }

    // Migrar o schema
    db.AutoMigrate(&models.Sample{}, &models.UserModel{})

    return db
}
