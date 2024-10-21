package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID        string  `gorm:"type:uuid;primaryKey"`
	OrderID   string  `gorm:"type:uuid"` // Relacionamento com o pedido
	ProductID string  `gorm:"type:uuid"` // Relacionamento com o produto
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"` // Quantidade do produto no pedido
	Price     float64 `json:"price"`    // Preço unitário do produto
	Total     float64 `json:"total"`    // Total do item (Quantity * Price)
}

// BeforeCreate: Gera um UUID para o Item do Pedido
func (item *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	if item.ID == "" {
		item.ID = uuid.New().String()
	}
	return nil
}
