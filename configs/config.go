package configs

import (
	"fmt"
	"log"

	"github.com/grupoG/csw24-grupoG-ticket-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Função para configurar e migrar o banco de dados
func SetupDatabase() *gorm.DB {
	host := "postgresdb"
	user := "postgres"
	password := "postgres"
	dbname := "mydb"
	port := "5432"

	// String de conexão
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	// Conectando ao banco de dados
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados:", err)
	}

	// Migrar o schema
	err = db.AutoMigrate(
		&models.Sample{},
	)
	if err != nil {
		log.Fatal("Falha ao migrar o schema:", err)
	}

	// Inserir dados iniciais
	createInitialData(db)

	return db
}

// Função para criar registros iniciais
func createInitialData(db *gorm.DB) {
	// Criando Samples
	sample1 := models.Sample{ID: 1, Name: "Sample A", Email: "Description A"}
	sample2 := models.Sample{ID: 2, Name: "Sample B", Email: "Description B"}
	db.Create(&sample1)
	db.Create(&sample2)

	fmt.Println("Dados iniciais criados com sucesso!")
}
