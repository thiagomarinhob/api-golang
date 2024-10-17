package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID              string          `gorm:"type:uuid;primaryKey"`
	CartID          string          `json:"cart_id"`
	Cart            Cart            `gorm:"foreignKey:CartID"`
	EstablishmentID string          `json:"establishment_id"`
	Establishment   Establishment   `gorm:"foreignKey:EstablishmentID"`
	Timeline        []OrderTimeline `gorm:"foreignKey:OrderID"`
}

func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	order.ID = uuid.New().String() // Gerar o UUID antes de criar o produto
	return
}
