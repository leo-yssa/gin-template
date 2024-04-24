package models

import (
	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	Kind string `gorm:"type:enum('email', 'oauth');not null" binding:"required"`
	Completed bool `gorm:"default:false"`
	Retrial uint8 `gorm:"default:0"`
	User User `gorm:"foreignKey:Id"`
}

