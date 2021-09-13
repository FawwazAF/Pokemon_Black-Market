package controller

import (
	"net/http"
	"project/pokemon/database"
	"project/pokemon/models"

	"github.com/labstack/echo"
)

func PostAddSeller(c echo.Context) error {
	//Bind data and save
	seller := models.Seller{}
	c.Bind(&seller)

	saved_seller, err := database.AddSeller(seller)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, saved_seller)
}

func GetSellerList(c echo.Context) error {
	//Get List of Transaction from Database
	seller, err := database.GetSellerList()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get seller list",
		})
	}
	if len(seller) == 0 {
		return c.JSON(http.StatusInternalServerError, "there's no seller")
	}
	return c.JSON(http.StatusOK, seller)
}

// Add Transaction Record
func PostRecordSoldPokemon(c echo.Context) error {
	//Bind data
	transaction := models.Transaction{}
	c.Bind(&transaction)

	//Search Pokemon and Seller
	pokemon, err := database.GetPokemonById(transaction.PokemonID) //Func in database/pokemon.go
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get pokemon from database",
		})
	}
	seller, err := database.GetSellerById(transaction.SellerID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get pokemon from database",
		})
	}

	//Save data
	new_transaction := models.Transaction{
		Seller:  seller,
		Pokemon: pokemon,
	}
	c.Bind(&new_transaction)
	saved_transaction, err := database.RecordTransaction(new_transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, saved_transaction)
}

func GetAllRecordedTransaction(c echo.Context) error {
	//Get List of Transaction from Database
	transactions, err := database.GetTransactionList()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get transaction list",
		})
	}
	if len(transactions) == 0 {
		return c.JSON(http.StatusInternalServerError, "there's no transaction")
	}
	return c.JSON(http.StatusOK, transactions)
}
