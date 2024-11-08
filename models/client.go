package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID        string  	`gorm:"type:uuid;primaryKey"`
	Name      string  	`json:"name"`
	Email     string  	`json:"email"`
	Addresses []Address `gorm:"foreignKey:ClientID"`
	Orders    []Order 	`gorm:"foreignKey:ClientID"`
}

func (client *Client) BeforeCreate(tx *gorm.DB) (err error) {
	if client.ID == "" {
		client.ID = uuid.New().String() // Gerar o UUID antes de criar o produto
	}
	return
}
