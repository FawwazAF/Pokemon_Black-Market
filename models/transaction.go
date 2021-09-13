package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	PokemonID uint    `json:"pokemon_id" form:"pokemon_id"`
	SellerID  uint    `json:"seller_id" form:"seller_id"`
	Price     int     `json:"price" form:"price"`
	Quantity  int     `json:"quantity" form:"quantity"`
	Pokemon   Pokemon `gorm:"foreignKey:PokemonID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	Seller    Seller  `gorm:"foreignKey:SellerID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
}
