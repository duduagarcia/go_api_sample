package models

import (
	"log"
	"time"

	"github.com/grupoG/csw24-grupoG-ticket-gin/models"
	"gorm.io/gorm"
)

// Função principal que executa todas as seeds
func RunSeeds(db *gorm.DB) {
	log.Println("Iniciando as seeds...")

	if err := createTenants(db); err != nil {
		log.Fatal("Erro ao criar tenants:", err)
	}
	if err := createUsers(db); err != nil {
		log.Fatal("Erro ao criar users:", err)
	}
	if err := createEvents(db); err != nil {
		log.Fatal("Erro ao criar events:", err)
	}
	if err := createTickets(db); err != nil {
		log.Fatal("Erro ao criar tickets:", err)
	}
	if err := createTransactions(db); err != nil {
		log.Fatal("Erro ao criar transactions:", err)
	}
	if err := createNotificationPreferences(db); err != nil {
		log.Fatal("Erro ao criar notification preferences:", err)
	}

	log.Println("Seeds finalizadas com sucesso!")
}

// Seed para Tenants
func createTenants(db *gorm.DB) error {
	tenants := []models.Tenant{
		{Name: "Tenant A", ContactInfo: "contact@tenanta.com", SpecificConfigurations: "Config A"},
		{Name: "Tenant B", ContactInfo: "contact@tenantb.com", SpecificConfigurations: "Config B"},
	}

	for _, tenant := range tenants {
		if err := db.FirstOrCreate(&tenant, models.Tenant{Name: tenant.Name}).Error; err != nil {
			return err
		}
	}
	return nil
}

// Seed para Users
func createUsers(db *gorm.DB) error {
	users := []models.User{
		{Name: "Alice", Email: "alice@example.com", TenantID: 1},
		{Name: "Bob", Email: "bob@example.com", TenantID: 2},
	}

	for _, user := range users {
		if err := db.FirstOrCreate(&user, models.User{Email: user.Email}).Error; err != nil {
			return err
		}
	}
	return nil
}

// Seed para Events
func createEvents(db *gorm.DB) error {
	events := []models.Event{
		{EventName: "Evento 1", EventDateTime: time.Now(), TenantID: 1},
		{EventName: "Evento 2", EventDateTime: time.Now().Add(24 * time.Hour), TenantID: 2},
	}

	for _, event := range events {
		if err := db.FirstOrCreate(&event, models.Event{EventName: event.EventName}).Error; err != nil {
			return err
		}
	}
	return nil
}

// Seed para Tickets
func createTickets(db *gorm.DB) error {
	tickets := []models.Ticket{
		{EventID: 1, TenantID: 1, OriginalPrice: 100, SellerID: 1, VerificationCode: "ABC123", Status: "Available"},
		{EventID: 2, TenantID: 2, OriginalPrice: 150, SellerID: 2, VerificationCode: "XYZ789", Status: "Available"},
	}

	for _, ticket := range tickets {
		if err := db.FirstOrCreate(&ticket, models.Ticket{VerificationCode: ticket.VerificationCode}).Error; err != nil {
			return err
		}
	}
	return nil
}

// Seed para Transactions
func createTransactions(db *gorm.DB) error {
	transactions := []models.Transaction{
		{TenantID: 1, BuyerID: 2, TicketID: 1, SalePrice: 95, TransactionDate: time.Now(), TransactionStatus: "Completed"},
		{TenantID: 2, BuyerID: 1, TicketID: 2, SalePrice: 145, TransactionDate: time.Now(), TransactionStatus: "Pending"},
	}

	for _, transaction := range transactions {
		if err := db.FirstOrCreate(&transaction, models.Transaction{TicketID: transaction.TicketID}).Error; err != nil {
			return err
		}
	}
	return nil
}

// Seed para NotificationPreferences
func createNotificationPreferences(db *gorm.DB) error {
	preferences := []models.NotificationPreferences{
		{UserID: 1, ReceiveEmails: true},
		{UserID: 2, ReceiveEmails: false},
	}

	for _, pref := range preferences {
		if err := db.FirstOrCreate(&pref, models.NotificationPreferences{UserID: pref.UserID}).Error; err != nil {
			return err
		}
	}
	return nil
}
