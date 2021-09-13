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

func RecordTransaction(transaction models.Transaction) (models.Transaction, error) {
	if err := config.DB.Save(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func GetTransactionList() ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := config.DB.Find(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil
}
