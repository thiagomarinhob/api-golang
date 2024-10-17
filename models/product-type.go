package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductType struct {
	gorm.Model
	ID       string    `gorm:"type:uuid;primaryKey"`
	Name     string    `json:"name"`
	Products []Product `gorm:"foreignKey:ProductTypeID"`
}

func (productType *ProductType) BeforeCreate(tx *gorm.DB) (err error) {
	productType.ID = uuid.New().String() // Gerar o UUID antes de criar o produto
	return
}
