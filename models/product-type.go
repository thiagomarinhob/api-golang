package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductType struct {
	gorm.Model
	ID              string        `gorm:"type:uuid;primaryKey"`
	Name            string        `json:"name"`
	Products        []Product     `gorm:"foreignKey:ProductTypeID"`
	EstablishmentID string        `json:"establishment_id"`
	Establishment   Establishment `gorm:"foreignKey:EstablishmentID"`
}

func (productType *ProductType) BeforeCreate(tx *gorm.DB) (err error) {
	if productType.ID == "" {
		productType.ID = uuid.New().String() // Gerar o UUID antes de criar o produto
	}
	return
}
