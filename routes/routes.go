package routes

import (
	"project/pokemon/controller"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {

	e.GET("/pokemons/:pokemon_id", controller.GetPokemonInPokedex)
	e.POST("/pokemons/:pokemon_id", controller.PostCreatePokemonInDatabase)
	e.PUT("/pokemons/:pokemon_id", controller.PutEditStockPokemon)
	e.DELETE("/pokemons/:pokemon_id", controller.DeletePokemonInDatabase)
	e.GET("/pokemons", controller.SearchAskedPokemon)

	e.GET("/sellers", controller.GetSellerList)
	e.POST("/sellers", controller.PostAddSeller)
	e.POST("/records", controller.PostRecordSoldPokemon)
	e.GET("/records", controller.GetAllRecordedTransaction)
	e.PUT("/records/:record_id", controller.PutEditRecalimedTransaction)
}
