package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID              string         `gorm:"type:uuid;primaryKey"`
  ClientID 				*string 			 `json:"client_id,omitempty" gorm:"type:uuid;default:null"`
	EstablishmentID string         `gorm:"type:uuid"`
	Status          string         `json:"status"`
	Items           []OrderItem    `gorm:"foreignKey:OrderID"`
	TotalAmount     float64         `json:"total_amount"`
	HistoryLog      []OrderHistory `gorm:"type:jsonb"` // Um campo JSONB para armazenar o histórico do pedido
}

// Definição da estrutura do histórico de pedido
type OrderHistory struct {
	Status    string    `json:"status"`    // O status do pedido em determinado momento
	Timestamp time.Time `json:"timestamp"` // Quando a mudança ocorreu
}

// BeforeCreate: Gera um UUID para o Pedido
func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if order.ID == "" {
		order.ID = uuid.New().String()
	}
	return nil
}
