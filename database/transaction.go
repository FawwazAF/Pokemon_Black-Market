package database

import (
	"project/pokemon/config"
	"project/pokemon/models"
)

func AddSeller(seller models.Seller) (models.Seller, error) {
	if err := config.DB.Save(&seller).Error; err != nil {
		return seller, err
	}
	return seller, nil
}

func GetSellerList() ([]models.Seller, error) {
	var seller []models.Seller
	if err := config.DB.Find(&seller).Error; err != nil {
		return seller, err
	}
	return seller, nil
}

func GetSellerById(seller_id uint) (models.Seller, error) {
	var seller models.Seller
	if err := config.DB.Find(&seller, "id=?", seller_id).Error; err != nil {
		return seller, err
	}
	return seller, nil
}

func RecordTransaction(transaction models.Transaction, pokemon models.Pokemon) (models.Transaction, error) {
	if err := config.DB.Save(&transaction).Error; err != nil {
		return transaction, err
	}
	//Auto Update Stock
	pokemon.Stock -= transaction.Quantity
	if err := config.DB.Save(&pokemon).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func GetListofTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := config.DB.Preload("Pokemon").Preload("Seller").Find(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil
}

func GetTransactionById(transaction_id int) (models.Transaction, error) {
	var transaction models.Transaction
	if err := config.DB.Preload("Pokemon").Preload("Seller").First(&transaction, "id=?", transaction_id).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func EditClaimedWarrantStock(transaction models.Transaction) (models.Transaction, error) {
	var pokemon models.Pokemon
	if err := config.DB.Find(&pokemon, "id = ?", transaction.PokemonID).Error; err != nil {
		return transaction, err
	}

	//Update Stock
	pokemon.Stock += transaction.Quantity
	if err := config.DB.Save(&pokemon).Error; err != nil {
		return transaction, err
	}

	//Edit Transaction Status
	transaction.Status = "reclaimed"
	transaction.Pokemon = pokemon
	if err := config.DB.Save(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}
