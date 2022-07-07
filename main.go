package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", showPokemon)
	e.GET("/getdatumByName", showDatumByName)
	e.GET("/getdatumById", showDatumById)
	e.Logger.Fatal(e.Start(":1323")) //e.loggerがe.post,e.getより先に書かれているとmessage not foundとなる。なぜか。
}

//グローバル変数
var pokemons []Pokemons

//typeたち
type Response struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
type Pokemons struct {
	Name string `json: "name"`
	Url  string `json: "url"`
	Id   int    `json: "id"`
}

//関数たち
func showPokemon(c echo.Context) error {
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

	var response Response
	var responses []Response

	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}
	responses = append(responses, response)
	for i := 0; i <= 0; i++ {
		for j := 0; j < 20; j++ {
			pokemons = append(pokemons, Pokemons{Name: responses[i].Results[j].Name, Url: responses[i].Results[j].URL, Id: i*20 + j + 1})
		}
	}
	/////////////////////////////
	for i := 1; i < 57; i++ {
		var response Response
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
			pokemons = append(pokemons, Pokemons{Name: responses[i].Results[j].Name, Url: responses[i].Results[j].URL, Id: i*20 + j + 1})
		}
	}
	/////////////////////////////

	return c.JSON(http.StatusOK, pokemons)
}

// func showPokemon(c echo.Context) error {
// 	url := "https://pokeapi.co/api/v2/pokemon"
// 	url2 := "https://pokeapi.co/api/v2/pokemon?offset=20\u0026limit=20"
// 	//以下他のapiから情報を持ってきてデータをjson形式に加工（もし情報を取ってくる先のurlの数が増えたら、その分だけ行数が増える）
// 	res, err := http.Get(url)
// 	if err != nil {
// 		panic(err)
// 	}
// 	res2, err := http.Get(url2)
// 	if err != nil {
// 		panic(err)
// 	}
// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	body2, err := ioutil.ReadAll(res2.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer res.Body.Close()
// 	defer res2.Body.Close()

// 	var response, response2 Response
// 	var responses []Response

// 	if err := json.Unmarshal(body, &response); err != nil {
// 		panic(err)
// 	}
// 	if err := json.Unmarshal(body2, &response2); err != nil {
// 		panic(err)
// 	}
// 	responses = append(responses, response, response2)
// 	for i := 0; i <= 1; i++ {
// 		for j := 0; j < 20; j++ {
// 			pokemons = append(pokemons, Pokemons{Name: responses[i].Results[j].Name, Url: responses[i].Results[j].URL, Id: i*20 + j + 1})
// 		}
// 	}
// 	return c.JSON(http.StatusOK, pokemons)
// }

func showDatumByName(c echo.Context) error {
	name := c.FormValue("name")
	for i := 0; i <= 40; i++ {
		if name == pokemons[i].Name {
			return c.JSON(http.StatusOK, pokemons[i].Id)
		}
	}
	return c.JSON(http.StatusOK, "missing name")
}
func showDatumById(c echo.Context) error {
	fmt.Println("done")
	name, _ := strconv.Atoi(c.FormValue("number"))
	for i := 0; i < 40; i++ {
		fmt.Println("name: ", name)
		fmt.Println("pokemons[i].Id: ", pokemons[i].Id)
		if name == pokemons[i].Id {
			return c.JSON(http.StatusOK, pokemons[i].Name)
		}
	}
	return c.JSON(http.StatusOK, "missing id")
}

//今できていること。。ポケモンの名前をhtmlページから検索してidを返す
