package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project/pokemon/database"
	"project/pokemon/models"
	"strconv"

	"github.com/labstack/echo"
)

// Consume pokedex API to check pokemon
func GetPokemonInPokedex(c echo.Context) error {
	//Reading path
	pokemon_id, err := strconv.Atoi(c.Param("pokemon_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid pokemon id",
		})
	}

	//Consuming Pokedex API
	pokedex_url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", pokemon_id)
	response, err := http.Get(pokedex_url)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}
	response_data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var pokemon models.Pokedex
	json.Unmarshal(response_data, &pokemon)

	return c.JSON(http.StatusOK, pokemon)
}

// Add Pokemon to Database
func PostCreatePokemonInDatabase(c echo.Context) error {
	//Reading path
	pokemon_id, err := strconv.Atoi(c.Param("pokemon_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid pokemon id",
		})
	}

	//Consuming Pokedex API
	pokedex_url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", pokemon_id)
	response, err := http.Get(pokedex_url)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}
	response_data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var pokedex models.Pokedex
	json.Unmarshal(response_data, &pokedex)

	//Binding request body
	pokemon := models.Pokemon{
		ID:   pokedex.ID,
		Name: pokedex.Name,
	}
	c.Bind(&pokemon)

	//Add Pokemon
	new_pokemon, err := database.AddPokemon(pokemon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, new_pokemon)
}

//Edit Existing Pokemon
func PutEditStockPokemon(c echo.Context) error {
	//Reading path
	pokemon_id, err := strconv.Atoi(c.Param("pokemon_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid pokemon id",
		})
	}

	//Get Pokemon from Database
	pokemon, err := database.GetPokemonFromDatabase(pokemon_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot get pokemon from database",
		})
	}

	//Edit Pokemon
	c.Bind(&pokemon)

	editted_pokemon, err := database.EditPokemon(pokemon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot edit pokemon's stock",
		})
	}
	return c.JSON(http.StatusOK, editted_pokemon)
}

//Delete Pokemon from Database
func DeletePokemonInDatabase(c echo.Context) error {
	//Reading path
	pokemon_id, err := strconv.Atoi(c.Param("pokemon_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid pokemon id",
		})
	}

	//Delete Pokemon
	deletted_pokemon, err := database.DeletePokemon(pokemon_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, deletted_pokemon)
}

// Get Asked Pokemon by customer
func SearchAskedPokemon(c echo.Context) error {
	//Reading path
	pokemon_name := c.QueryParam("pokemon_name")

	//Search Pokemon
	pokemons, err := database.SearchPokemon(pokemon_name)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, pokemons)
}
