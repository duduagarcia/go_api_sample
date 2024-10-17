package models

import (
	"gorm.io/gorm"
)

type TenantModel struct {
	gorm.Model                        
	Name                  string      `gorm:"size:100;not null"` 
	ContactInfo           string      `gorm:"size:255"`          
	SpecificConfigurations string     `gorm:"size:255"`          
	Users                 []UserModel      `gorm:"foreignKey:TenantID"` // One-to-many relationship with User
	Events                []EventModel     `gorm:"foreignKey:TenantID"` // One-to-many relationship with Event
	Tickets               []TicketModel    `gorm:"foreignKey:TenantID"` // One-to-many relationship with Ticket
	Transactions          []TransactionModel `gorm:"foreignKey:TenantID"` // One-to-many relationship with Transaction
}


