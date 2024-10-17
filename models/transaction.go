package models

import (
	"gorm.io/gorm"
)

type TransactionModel struct {
	gorm.Model              
	TenantID       uint     `gorm:"not null"`            // FK to Tenant
	BuyerID        uint     `gorm:"not null"`            // FK to User (buyer)
	TicketID       uint     `gorm:"not null;unique"`     // FK to Ticket (unique)
	SalePrice      float64  `gorm:"not null"`            
	TransactionDate time.Time `gorm:"not null"`          
	TransactionStatus string `gorm:"size:50;not null"`   // Transaction status (pending, completed, canceled)
}

