package models

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model                 
	EventID       uint          `gorm:"not null"`           // FK to Event
	TenantID      uint          `gorm:"not null"`           // FK to Tenant
	OriginalPrice float64       `gorm:"not null"`           
	SellerID      uint          `gorm:"not null"`           // FK to User (seller)
	VerificationCode string     `gorm:"size:100;not null"`  
	Status        string        `gorm:"size:50;not null"`   
	Transaction   TransactionModel   `gorm:"foreignKey:TicketID"` // One-to-one relationship with Transaction
}

