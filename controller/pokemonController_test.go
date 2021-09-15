package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"project/pokemon/config"
	"project/pokemon/models"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetPokedex(t *testing.T) {
	//Consuming Pokedex API
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/1")
	if err != nil {
		t.Error(err)
	}
	response_data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var pokemon models.Pokedex
	json.Unmarshal(response_data, &pokemon)

	//setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/pokemons/:pokemone_id")
	context.SetParamNames("pokemon_id")
	context.SetParamValues("1")

	//Call Func
	GetPokemonInPokedex(context)

	//unmarshal response
	var test_response models.Pokedex
	req_body2 := res.Body.String()
	json.Unmarshal([]byte(req_body2), &test_response)

	//Test expected and actual result
	t.Run("GET /pokemons/:pokemone_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, pokemon.ID, test_response.ID)
		assert.Equal(t, pokemon.Name, test_response.Name)
	})
}

func TestPostCreatePokemon(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning and migrating data before test
	db.Migrator().DropTable(&models.Pokemon{})
	db.AutoMigrate(&models.Pokemon{})

	//Consuming Pokedex API
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/1")
	if err != nil {
		t.Error(err)
	}
	response_data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var pokemon models.Pokedex
	json.Unmarshal(response_data, &pokemon)

	//Setting req body
	req_body, _ := json.Marshal(map[string]interface{}{
		"stock": 50,
	})

	//setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(req_body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/pokemons/:pokemone_id")
	context.SetParamNames("pokemon_id")
	context.SetParamValues("1")

	//Call Func
	PostCreatePokemonInDatabase(context)

	//unmarshal response
	var test_response models.Pokemon
	res_body := res.Body.String()
	json.Unmarshal([]byte(res_body), &test_response)

	//Test expected and actual result
	t.Run("POST /pokemons/:pokemone_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, pokemon.ID, test_response.ID)
		assert.Equal(t, pokemon.Name, test_response.Name)
		assert.Equal(t, 50, test_response.Stock)
	})
}

func TestPutEditPokemonStock(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning and migrating data before test
	db.Migrator().DropTable(&models.Pokemon{})
	db.AutoMigrate(&models.Pokemon{})

	//Pokemon Dummy
	new_pokemon := models.Pokemon{
		ID:    uint(1),
		Name:  "pikachu",
		Stock: 5,
	}
	if err := db.Create(&new_pokemon).Error; err != nil {
		t.Error(err)
	}

	//Setting req body
	req_body, _ := json.Marshal(map[string]interface{}{
		"stock": 1,
	})

	//setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(req_body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/pokemons/:pokemone_id")
	context.SetParamNames("pokemon_id")
	context.SetParamValues("1")

	//Call Func
	PutEditStockPokemon(context)

	//unmarshal response
	var test_response models.Pokemon
	res_body := res.Body.String()
	json.Unmarshal([]byte(res_body), &test_response)

	//Test expected and actual result
	t.Run("PUT /pokemons/:pokemone_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), test_response.ID)
		assert.Equal(t, 1, test_response.Stock)
	})
}

func TestDeletePokemon(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning and migrating data before test
	db.Migrator().DropTable(&models.Pokemon{})
	db.AutoMigrate(&models.Pokemon{})

	//Pokemon Dummy
	new_pokemon := models.Pokemon{
		ID:    uint(1),
		Name:  "pikachu",
		Stock: 5,
	}
	if err := db.Create(&new_pokemon).Error; err != nil {
		t.Error(err)
	}

	//setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/pokemons/:pokemone_id")
	context.SetParamNames("pokemon_id")
	context.SetParamValues("1")

	//Call Func
	DeletePokemonInDatabase(context)

	//unmarshal response
	var test_response models.Pokemon
	res_body := res.Body.String()
	json.Unmarshal([]byte(res_body), &test_response)

	//Test expected and actual result
	t.Run("DELETE /pokemons/:pokemone_id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, uint(1), test_response.ID)
	})
}

func TestGetSearchPokemon(t *testing.T) {
	// create database connection
	db, err := config.ConfigTest()
	if err != nil {
		t.Error(err)
	}

	// cleaning and migrating data before test
	db.Migrator().DropTable(&models.Pokemon{})
	db.AutoMigrate(&models.Pokemon{})

	//Pokemon Dummy
	new_pokemon := models.Pokemon{
		ID:    uint(1),
		Name:  "bulba",
		Stock: 5,
	}
	if err := db.Create(&new_pokemon).Error; err != nil {
		t.Error(err)
	}
	new_pokemon2 := models.Pokemon{
		ID:    uint(2),
		Name:  "pikachu",
		Stock: 5,
	}
	if err := db.Create(&new_pokemon2).Error; err != nil {
		t.Error(err)
	}

	//setting controller
	e := echo.New()
	q := make(url.Values)
	q.Set("pokemon_name", "pikachu")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/pokemons")

	//Call Func
	SearchAskedPokemon(context)

	//unmarshal response
	var test_response []models.Pokemon
	req_body2 := res.Body.String()
	json.Unmarshal([]byte(req_body2), &test_response)

	//Test expected and actual result
	t.Run("GET /pokemons", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, new_pokemon2.ID, test_response[0].ID)
		assert.Equal(t, new_pokemon2.Name, test_response[0].Name)
	})
}
