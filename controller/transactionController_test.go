package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project/pokemon/config"
	"project/pokemon/models"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestPostAddSeller(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning and migrating data before test
	db.Migrator().DropTable(&models.Seller{})
	db.AutoMigrate(&models.Seller{})

	//Setting req body
	req_body, _ := json.Marshal(map[string]interface{}{
		"name": "minato",
	})

	//setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(req_body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/sellers")

	//Call Func
	PostAddSeller(context)

	//unmarshal response
	var test_response models.Seller
	req_body2 := res.Body.String()
	json.Unmarshal([]byte(req_body2), &test_response)

	//Test expected and actual result
	t.Run("POST /sellers", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "minato", test_response.Name)
	})
}

func TestGetSeller(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning and migrating data before test
	db.Migrator().DropTable(&models.Seller{})
	db.AutoMigrate(&models.Seller{})

	//Pokemon Dummy
	new_seller := models.Seller{
		Name: "minato",
	}
	if err := db.Create(&new_seller).Error; err != nil {
		t.Error(err)
	}

	//setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/sellers")

	//Call Func
	GetSellerList(context)

	//unmarshal response
	var test_response []models.Seller
	res_body := res.Body.String()
	json.Unmarshal([]byte(res_body), &test_response)

	//Test expected and actual result
	t.Run("GET /sellers", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, new_seller.Name, test_response[0].Name)
	})
}

func TestPostRecordTransaction(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning and migrating data before test
	db.Migrator().DropTable(&models.Seller{}, &models.Pokemon{}, &models.Transaction{})
	db.AutoMigrate(&models.Seller{}, &models.Pokemon{}, &models.Transaction{})

	//Data Dummy
	new_pokemon := models.Pokemon{
		Name:  "bulba",
		Stock: 5,
	}
	if err := db.Create(&new_pokemon).Error; err != nil {
		t.Error(err)
	}

	new_seller := models.Seller{
		Name: "minato",
	}
	if err := db.Create(&new_seller).Error; err != nil {
		t.Error(err)
	}

	//Setting req body
	req_body, _ := json.Marshal(map[string]interface{}{
		"pokemon_id":  1,
		"seller_id":   1,
		"quantity":    3,
		"total_price": 5000,
	})

	//setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(req_body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/records")

	//Call Func
	PostRecordSoldPokemon(context)

	//unmarshal response
	var test_response models.Transaction
	req_body2 := res.Body.String()
	json.Unmarshal([]byte(req_body2), &test_response)

	//Test expected and actual result
	t.Run("POST /records", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), test_response.PokemonID)
		assert.Equal(t, uint(1), test_response.SellerID)
		assert.Equal(t, 3, test_response.Quantity)
		assert.Equal(t, 5000, test_response.TotalPrice)
		assert.Equal(t, "Sold", test_response.Status)
	})
}

func TestGetListTransaction(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning and migrating data before test
	db.Migrator().DropTable(&models.Seller{}, &models.Pokemon{}, &models.Transaction{})
	db.AutoMigrate(&models.Seller{}, &models.Pokemon{}, &models.Transaction{})

	//Data Dummy
	new_pokemon := models.Pokemon{
		Name:  "bulba",
		Stock: 5,
	}
	if err := db.Create(&new_pokemon).Error; err != nil {
		t.Error(err)
	}

	new_seller := models.Seller{
		Name: "minato",
	}
	if err := db.Create(&new_seller).Error; err != nil {
		t.Error(err)
	}
	new_transaction := models.Transaction{
		PokemonID:  uint(1),
		SellerID:   uint(1),
		Quantity:   2,
		TotalPrice: 5000,
		Status:     "Sold",
	}
	if err := db.Create(&new_transaction).Error; err != nil {
		t.Error(err)
	}
	//setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/records")

	//Call Func
	GetAllRecordedTransaction(context)

	//unmarshal response
	var test_response []models.Transaction
	req_body2 := res.Body.String()
	json.Unmarshal([]byte(req_body2), &test_response)

	//Test expected and actual result
	t.Run("GET /records", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), test_response[0].PokemonID)
		assert.Equal(t, uint(1), test_response[0].SellerID)
		assert.Equal(t, 2, test_response[0].Quantity)
		assert.Equal(t, 5000, test_response[0].TotalPrice)
		assert.Equal(t, "sold", test_response[0].Status)
	})
}

func TestPutWarrantyClaim(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning and migrating data before test
	db.Migrator().DropTable(&models.Seller{}, &models.Pokemon{}, &models.Transaction{})
	db.AutoMigrate(&models.Seller{}, &models.Pokemon{}, &models.Transaction{})

	//Data Dummy
	new_pokemon := models.Pokemon{
		Name:  "bulba",
		Stock: 5,
	}
	if err := db.Create(&new_pokemon).Error; err != nil {
		t.Error(err)
	}

	new_seller := models.Seller{
		Name: "minato",
	}
	if err := db.Create(&new_seller).Error; err != nil {
		t.Error(err)
	}
	new_transaction := models.Transaction{
		PokemonID:  uint(1),
		SellerID:   uint(1),
		Quantity:   2,
		TotalPrice: 5000,
		Status:     "Sold",
	}
	if err := db.Create(&new_transaction).Error; err != nil {
		t.Error(err)
	}
	//setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/records/:record_id")
	context.SetParamNames("record_id")
	context.SetParamValues("1")

	//Call Func
	PutEditRecalimedTransaction(context)

	//unmarshal response
	var test_response models.Transaction
	req_body2 := res.Body.String()
	json.Unmarshal([]byte(req_body2), &test_response)

	//Test expected and actual result
	t.Run("PUT /records/:record_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), test_response.PokemonID)
		assert.Equal(t, uint(1), test_response.SellerID)
		assert.Equal(t, 2, test_response.Quantity)
		assert.Equal(t, 5000, test_response.TotalPrice)
		assert.Equal(t, "reclaimed", test_response.Status)
	})
}
