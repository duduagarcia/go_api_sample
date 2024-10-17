package models

import (
	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model                        
	Name                  string      `gorm:"size:100;not null"` 
	ContactInfo           string      `gorm:"size:255"`          
	SpecificConfigurations string     `gorm:"size:255"`          
	Users                 []User      `gorm:"foreignKey:TenantID"` // One-to-many relationship with User
	Events                []Event     `gorm:"foreignKey:TenantID"` // One-to-many relationship with Event
	Tickets               []Ticket    `gorm:"foreignKey:TenantID"` // One-to-many relationship with Ticket
	Transactions          []Transaction `gorm:"foreignKey:TenantID"` // One-to-many relationship with Transaction
}


