package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderStatus struct {
	gorm.Model
	ID        string          `gorm:"type:uuid;primaryKey"`
	Name      string          `json:"name"`
	Timelines []OrderTimeline `gorm:"foreignKey:StatusID"`
}

func (orderStatus *OrderStatus) BeforeCreate(tx *gorm.DB) (err error) {
	orderStatus.ID = uuid.New().String() // Gerar o UUID antes de criar o produto
	return
}
