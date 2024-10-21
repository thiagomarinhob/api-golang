package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductPhoto struct {
	gorm.Model
	ID        string  `gorm:"type:uuid;primaryKey"`
	ProductID string  `json:"product_id"`
	URL       string  `json:"url"`
	Product   Product `gorm:"foreignKey:ProductID"`
}

func (productPhoto *ProductPhoto) BeforeCreate(tx *gorm.DB) (err error) {
	if productPhoto.ID == "" {
		productPhoto.ID = uuid.New().String() // Gerar o UUID antes de criar o produto
	}
	return
}
