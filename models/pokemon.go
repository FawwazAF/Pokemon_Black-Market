package models

import (
	"time"

	"gorm.io/gorm"
)

type Pokemon struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `json:"name"`
	Stock     int    `json:"stock" form:"stock"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
