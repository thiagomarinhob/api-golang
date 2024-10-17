package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID        string  `gorm:"type:uuid;primaryKey"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	AddressID string  `json:"address_id"`
	Address   Address `gorm:"foreignKey:AddressID"`
	Carts     []Cart  `gorm:"foreignKey:CustomerID"`
}

func (client *Client) BeforeCreate(tx *gorm.DB) (err error) {
	client.ID = uuid.New().String() // Gerar o UUID antes de criar o produto
	return
}
