package models

import (
	"gorm.io/gorm"
)

type Seller struct {
	gorm.Model
	Name string `json:"name" form:"name"`
}
