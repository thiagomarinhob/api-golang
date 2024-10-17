package models

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	Token     string `gorm:"unique;not null"`
	UserID    uint   `gorm:"not null"`
	ExpiresAt int64  `gorm:"not null"`
}
