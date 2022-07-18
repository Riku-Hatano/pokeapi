package funcs

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ShowDatumByName(c echo.Context) error {
	name := c.FormValue("name")
	for i := 0; i <= 1140; i++ {
		if name == Pokemons[i].Name {
			return c.JSON(http.StatusOK, Pokemons[i].Id)
		}
	}
	return c.JSON(http.StatusOK, "missing name")
}
