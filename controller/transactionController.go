package controller

import (
	"net/http"
	"project/pokemon/database"
	"project/pokemon/models"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

//Set a Response
type TransactionResponse struct {
	ID         uint   `json:"id"`
	PokemonID  uint   `json:"pokemon_id"`
	SellerID   uint   `json:"seller_id"`
	TotalPrice int    `json:"total_price"`
	Quantity   int    `json:"quantity"`
	Status     string `json:"status"`
}

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
	transaction := models.Transaction{
		Status: "Sold",
	}
	c.Bind(&transaction)

	//Search Pokemon
	pokemon, err := database.GetPokemonFromDatabase(int(transaction.PokemonID)) //Func in database/pokemon.go
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get pokemon from database",
		})
	}

	//Checking Stock
	if transaction.Quantity > pokemon.Stock {
		return c.JSON(http.StatusBadRequest, "stock is empty")
	}

	//Save Transaction to Database
	saved_transaction, err := database.RecordTransaction(transaction, pokemon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response := TransactionResponse{
		ID:         saved_transaction.ID,
		PokemonID:  saved_transaction.PokemonID,
		SellerID:   saved_transaction.SellerID,
		TotalPrice: saved_transaction.TotalPrice,
		Quantity:   saved_transaction.Quantity,
		Status:     saved_transaction.Status,
	}
	return c.JSON(http.StatusOK, response)
}

//Get List of Transaction from Database
func GetAllRecordedTransaction(c echo.Context) error {
	transactions, err := database.GetListofTransaction()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get transaction list",
		})
	}
	//Handling error
	if len(transactions) == 0 {
		return c.JSON(http.StatusInternalServerError, "there's no transaction")
	}
	return c.JSON(http.StatusOK, transactions)
}

//Warranty system auto update stock
func PutEditRecalimedTransaction(c echo.Context) error {
	//Reading path
	transaction_id, err := strconv.Atoi(c.Param("record_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid transaction id",
		})
	}

	//Get Transaction by id
	transaction, err := database.GetTransactionById(transaction_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	//Get Warranty expired date and time
	warranty_time := 3 // by day
	current_time := time.Now()
	warranty_date := transaction.CreatedAt.Add(time.Hour * 24 * time.Duration(warranty_time))

	//Checking if warranty is still available
	if current_time.After(warranty_date) {
		return c.JSON(http.StatusBadRequest, "warranty is expired")
	}
	if transaction.Status != "sold" {
		return c.JSON(http.StatusBadRequest, "warranty already claimed")
	}

	//Update pokemon stock in database
	editted_transaction, err := database.EditClaimedWarrantStock(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot edit pokemon's stock",
		})
	}

	//Edit Response
	response := TransactionResponse{
		ID:         editted_transaction.ID,
		PokemonID:  editted_transaction.PokemonID,
		SellerID:   editted_transaction.SellerID,
		TotalPrice: editted_transaction.TotalPrice,
		Quantity:   editted_transaction.Quantity,
		Status:     editted_transaction.Status,
	}
	return c.JSON(http.StatusOK, response)
}
