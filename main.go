package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/get", showPokemon)
	e.GET("/getpokemon", getPokemon)
	e.GET("/after", after)
	e.Logger.Fatal(e.Start(":1323")) //e.loggerがe.post,e.getより先に書かれているとmessage not foundとなる。なぜか。
}

//ハンドラーを定義

type Status struct {
	Name string `json: "name"`
	Url  string `json: "url"`
}

var statuses []Status

var html string

func after(c echo.Context) error {
	fmt.Println(html)
	return c.JSON(http.StatusOK, html)
}

func getPokemon(c echo.Context) error {
	url := "https://pokeapi.co/api/v2/pokemon"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// buf := bytes.NewBuffer(body)
	// html := buf.String()

	// bytes, err := ioutil.ReadFile(url)
	// if err != nil {
	// 	panic(err)
	// }
	//json decode
	var statuses []Status
	if err := json.Unmarshal(body, &statuses); err != nil {
		fmt.Println("done second")
		panic(err)
	}
	//show decoded data
	for _, p := range statuses {
		fmt.Println(p.Name, p.Url)
	}
	return c.JSON(http.StatusOK, statuses)
}

/////////////////////////////////////////////
func showPokemon(c echo.Context) error {
	url := "https://pokeapi.co/api/v2/pokemon"
	url2 := "https://pokeapi.co/api/v2/pokemon?offset=20\u0026limit=20"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	res2, err := http.Get(url2)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	body2, err := ioutil.ReadAll(res2.Body)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	defer res2.Body.Close()
	var response Response
	var response2 Response
	var responses []Response
	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body2, &response2); err != nil {
		panic(err)
	}
	responses = append(responses, response, response2)
	return c.JSON(http.StatusOK, responses)
}

type Response struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
