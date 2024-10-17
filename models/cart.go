package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID              string        `gorm:"type:uuid;primaryKey"`
	CustomerID      uint          `json:"customer_id"`
	EstablishmentID uint          `json:"establishment_id"`
	Establishment   Establishment `gorm:"foreignKey:EstablishmentID"`
	Client          Client        `gorm:"foreignKey:CustomerID"`
	Items           []CartItem    `gorm:"foreignKey:CartID"`
}

func (cart *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	cart.ID = uuid.New().String() // Gerar o UUID antes de criar o produto
	return
}
