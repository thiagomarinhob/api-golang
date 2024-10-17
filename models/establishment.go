package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Establishment struct {
	gorm.Model
	ID          string `gorm:"type:uuid;primaryKey"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	AdminUserID string `json:"admin_user_id"`
	AdminUser   User   `gorm:"foreignKey:AdminUserID"`
}

func (establishment *Establishment) BeforeCreate(tx *gorm.DB) (err error) {
	establishment.ID = uuid.New().String() // Gerar o UUID antes de criar o produto
	return
}
