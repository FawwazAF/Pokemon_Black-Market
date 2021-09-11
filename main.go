package main

import (
	"fmt"
	"project/pokemon/config"
	"project/pokemon/routes"

	"github.com/labstack/echo"
)

func main() {
	//deploy2
	e := echo.New()
	config.InitDb()
	config.InitPort()
	routes.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))

}
