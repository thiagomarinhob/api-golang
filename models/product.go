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
	Price           string         `json:"price"`
	Photos          []ProductPhoto `gorm:"foreignKey:ProductID"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if product.ID == "" {
		product.ID = uuid.New().String() // Gerar o UUID se ele não existir
	}
	return nil
}
