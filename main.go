package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"pokemonApi2/funcs"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/imgs", "./imgs")
	e.Static("/css", "./css")
	e.Static("/js", "./js")
	e.GET("/", funcs.ShowPokemon)
	e.GET("/getdatumByName", funcs.ShowDatumByName)
	e.GET("/getdatumById", funcs.ShowDatumById)
	e.Logger.Fatal(e.Start(":1323")) //e.loggerがe.post,e.getより先に書かれているとmessage not foundとなる。なぜか。
}
