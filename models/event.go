package models

import (
	"gorm.io/gorm"
)

type EventModel struct {
	gorm.Model         
	TenantID    uint   `gorm:"not null"`              
	EventName   string `gorm:"size:100;not null"`      
	Type        string `gorm:"size:50;not null"`       
	Location    string `gorm:"size:255"`               
	EventDateTime time.Time `gorm:"not null"`          // Date and time of the event
	Tickets     []TicketModel `gorm:"foreignKey:EventID"`   // One-to-many relationship with Ticket
}

