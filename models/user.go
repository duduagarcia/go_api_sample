package models

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model          
	TenantID    uint    `gorm:"not null"`               
	Name        string  `gorm:"size:100;not null"`      
	Email       string  `gorm:"size:100;unique;not null"` 
	Tickets     []TicketModel `gorm:"foreignKey:SellerID"`   // One-to-many relationship with Ticket as seller
	Transactions []TransactionModel `gorm:"foreignKey:BuyerID"` // One-to-many relationship with Transaction as buyer
	Preferences NotificationPreferencesModel `gorm:"foreignKey:UserID"` // One-to-one relationship with NotificationPreferences
}

