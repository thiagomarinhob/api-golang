package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderTimeline struct {
	gorm.Model
	ID       string      `gorm:"type:uuid;primaryKey"`
	OrderID  string      `json:"order_id"`
	StatusID string      `json:"status_id"`
	Status   OrderStatus `gorm:"foreignKey:StatusID"`
}

func (orderTimeline *OrderTimeline) BeforeCreate(tx *gorm.DB) (err error) {
	orderTimeline.ID = uuid.New().String() // Gerar o UUID antes de criar o produto
	return
}
