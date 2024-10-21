package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	ID       string   `gorm:"type:uuid;primaryKey"`
	Street   string   `json:"street"`
	City     string   `json:"city"`
	State    string   `json:"state"`
	Number   string   `json:"number"`
	District string   `json:"district"`
	Clients  []Client `gorm:"foreignKey:AddressID"`
}

func (address *Address) BeforeCreate(tx *gorm.DB) (err error) {
	if address.ID == "" {
		address.ID = uuid.New().String() // Gerar o UUID antes de criar o produto
	}
	return
}
