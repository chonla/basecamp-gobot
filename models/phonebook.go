package models

import "gorm.io/gorm"

// Phonebook model
type Phonebook struct {
	gorm.Model
	userID       int64  `gorm:"column:user_id; uniqueIndex"`
	phoneNumbers string `gorm:"column:phone_numbers"`
}
