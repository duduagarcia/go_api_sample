package entities

import (
	"gorm.io/gorm"
)

// UserModel is the GORM model for the User entity, including database annotations.
type UserModel struct {
	gorm.Model        // Adds fields like ID, CreatedAt, UpdatedAt, and DeletedAt
	Name       string `gorm:"size:100;not null"`        // Name with a size limit and "not null" constraint
	Email      string `gorm:"size:100;unique;not null"` // Email must be unique and "not null"
	Age        int    // Age of the user
}
