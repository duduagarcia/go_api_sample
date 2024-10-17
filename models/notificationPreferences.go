package models

import (
	"gorm.io/gorm"
)

type NotificationPreferencesModel struct {
	gorm.Model        
	UserID    uint    `gorm:"not null;unique"`         // FK to User (unique)
	ReceiveEmails bool `gorm:"not null"`               // Preference to receive emails
}

