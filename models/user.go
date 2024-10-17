package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model          
	TenantID    uint    `gorm:"not null"`               
	Name        string  `gorm:"size:100;not null"`      
	Email       string  `gorm:"size:100;unique;not null"` 
	Tickets     []Ticket `gorm:"foreignKey:SellerID"`   // One-to-many relationship with Ticket as seller
	Transactions []Transaction `gorm:"foreignKey:BuyerID"` // One-to-many relationship with Transaction as buyer
	Preferences NotificationPreferences `gorm:"foreignKey:UserID"` // One-to-one relationship with NotificationPreferences
}

