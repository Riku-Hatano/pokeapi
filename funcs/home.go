package funcs

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pokemonApi2/structs"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

//グローバル変数
var Pokemons []structs.Pokemons
var Statses []structs.ResponseStats

func ShowPokemon(c echo.Context) error {
	url := "https://pokeapi.co/api/v2/pokemon"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var response structs.Response
	var responses []structs.Response

	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}
	responses = append(responses, response)
	for i := 0; i <= 0; i++ {
		for j := 0; j < 20; j++ {
			Pokemons = append(Pokemons, structs.Pokemons{Name: responses[i].Results[j].Name, Url: responses[i].Results[j].URL, Id: i*20 + j + 1})
		}
	}
	//以下繰り返し処理でポケモンの情報を追加
	for i := 1; i < 57; i++ {
		var response structs.Response
		url := "https://pokeapi.co/api/v2/pokemon?offset=" + strconv.Itoa(i*20) + "\u0026limit=" + strconv.Itoa(i*20)
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(res.Body)
		if err := json.Unmarshal(body, &response); err != nil {
			panic(err)
		}
		responses = append(responses, response)
		for j := 0; j < 20; j++ {
			Pokemons = append(Pokemons, structs.Pokemons{Name: responses[i].Results[j].Name, Url: responses[i].Results[j].URL, Id: i*20 + j + 1})
		}
	}
	//繰り返し処理のポケモンの情報の追加終わり
	// return c.JSON(http.StatusOK, Pokemons)
	w := c.Response()
	t, _ := template.ParseFiles("main.html")
	return t.Execute(w, "")
}
