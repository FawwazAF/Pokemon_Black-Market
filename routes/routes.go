package routes

import (
	"project/pokemon/constant"
	"project/pokemon/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	e.GET("/pokemons/:pokemon_id", controller.GetPokemonInPokedex)
	e.POST("/pokemons/:pokemon_id", controller.PostCreatePokemonInDatabase)
	e.PUT("/pokemons/:pokemon_id", controller.PutEditStockPokemon)
	e.DELETE("/pokemons/:pokemon_id", controller.DeletePokemonInDatabase)
	e.GET("/pokemons", controller.SearchAskedPokemon)

	//Login
	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))

}
