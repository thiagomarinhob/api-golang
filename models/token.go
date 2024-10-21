package models

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	Token     string `gorm:"unique;not null"`
	UserID    string `gorm:"not null"`
	ExpiresAt int64  `gorm:"not null"`
}
