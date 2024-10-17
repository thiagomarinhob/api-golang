package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID              string         `gorm:"type:uuid;primaryKey"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	ProductTypeID   string         `json:"product_type_id"`
	ProductType     ProductType    `gorm:"foreignKey:ProductTypeID"`
	EstablishmentID string         `json:"establishment_id"`
	Establishment   Establishment  `gorm:"foreignKey:EstablishmentID"`
	Prices          string         `json:"price"`
	Photos          []ProductPhoto `gorm:"foreignKey:ProductID"`
	CartItems       []CartItem     `gorm:"foreignKey:ProductID"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.ID = uuid.New().String() // Gerar o UUID antes de criar o produto
	return
}
