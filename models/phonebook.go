package models

import "gorm.io/gorm"

// Phonebook model
type Phonebook struct {
	gorm.Model
	userID       int64 `gorm:"uniqueIndex"`
	phoneNumbers string
}
