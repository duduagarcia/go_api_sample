package configs

import (
	"fmt"
	"log"
	"time"

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
	port := "5432" // Porta padrão do Postgres

	// String de conexão
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	// Conectando ao banco de dados
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados:", err)
	}

	// Migrar o schema
	err = db.AutoMigrate(
		&models.Tenant{}, 
		&models.User{}, 
		&models.Event{}, 
		&models.Ticket{}, 
		&models.Transaction{}, 
		&models.NotificationPreferences{},
	)
	if err != nil {
		log.Fatal("Falha ao migrar o schema:", err)
	}

	// Inserir dados iniciais
	createInitialData(db)
	// seeds.RunSeeds(db)

	return db
}

// Função para criar registros iniciais
func createInitialData(db *gorm.DB) {
	// Criando Tenants
	tenant1 := models.Tenant{Name: "Tenant A", ContactInfo: "contact@tenanta.com", SpecificConfigurations: "Config A"}
	tenant2 := models.Tenant{Name: "Tenant B", ContactInfo: "contact@tenantb.com", SpecificConfigurations: "Config B"}
	db.Create(&tenant1)
	db.Create(&tenant2)

	// Criando Users
	user1 := models.User{Name: "Alice", Email: "alice@example.com", TenantID: tenant1.ID}
	user2 := models.User{Name: "Bob", Email: "bob@example.com", TenantID: tenant2.ID}
	db.Create(&user1)
	db.Create(&user2)

	// Criando Events
	event1 := models.Event{TenantID: tenant1.ID, EventName: "Conference", Type: "Business", Location: "New York", EventDateTime: time.Now().AddDate(0, 1, 0)}
	event2 := models.Event{TenantID: tenant2.ID, EventName: "Music Festival", Type: "Entertainment", Location: "Los Angeles", EventDateTime: time.Now().AddDate(0, 2, 0)}
	db.Create(&event1)
	db.Create(&event2)

	// Criando Tickets
	ticket1 := models.Ticket{EventID: event1.ID, TenantID: tenant1.ID, OriginalPrice: 100.00, SellerID: user1.ID, VerificationCode: "ABC123", Status: "Available"}
	ticket2 := models.Ticket{EventID: event2.ID, TenantID: tenant2.ID, OriginalPrice: 150.00, SellerID: user2.ID, VerificationCode: "XYZ789", Status: "Available"}
	db.Create(&ticket1)
	db.Create(&ticket2)

	// Criando Transactions
	transaction1 := models.Transaction{TenantID: tenant1.ID, BuyerID: user2.ID, TicketID: ticket1.ID, SalePrice: 95.00, TransactionDate: time.Now(), TransactionStatus: "Completed"}
	db.Create(&transaction1)

	// Criando NotificationPreferences
	preferences1 := models.NotificationPreferences{UserID: user1.ID, ReceiveEmails: true}
	preferences2 := models.NotificationPreferences{UserID: user2.ID, ReceiveEmails: false}
	db.Create(&preferences1)
	db.Create(&preferences2)

	fmt.Println("Dados iniciais criados com sucesso!")
}
