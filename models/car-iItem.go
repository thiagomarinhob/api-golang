package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	ID        string  `gorm:"type:uuid;primaryKey"`
	CartID    string  `json:"cart_id"`
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Product   Product `gorm:"foreignKey:ProductID"`
}

func (cartItem *CartItem) BeforeCreate(tx *gorm.DB) (err error) {
	cartItem.ID = uuid.New().String() // Gerar o UUID antes de criar o produto
	return
}
