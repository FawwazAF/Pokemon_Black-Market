package routes

import (
	"project/pokemon/constant"
	"project/pokemon/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	e.GET("/pokemons/:pokemon_id", controller.GetPokemonInPokedex)

	//Login
	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))

}
