package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "project/pokemon/database"
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
