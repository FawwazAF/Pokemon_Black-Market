package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	PokemonID  uint    `json:"pokemon_id" form:"pokemon_id"`
	SellerID   uint    `json:"seller_id" form:"seller_id"`
	TotalPrice int     `json:"total_price" form:"total_price"`
	Quantity   int     `json:"quantity" form:"quantity"`
	Status     string  `gorm:"type:enum('sold', 'reclaimed')" json:"status" form:"status"`
	Pokemon    Pokemon `gorm:"foreignKey:PokemonID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	Seller     Seller  `gorm:"foreignKey:SellerID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
}
