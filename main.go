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
	// Previous string `json: "previous"`
	// Results  []string `json: "results"`
}
type Response struct {
	Results []Status `json: "results"`
}

var statuses []Status

// func showPokemon(c echo.Context) error {
// 	fmt.Println("done by pokemon")
// 	return c.JSON(http.StatusOK, users)
// }
var html string

func showPokemon(c echo.Context) error {
	fmt.Println("done again")
	fmt.Println("done")
	var p Status
	if err := c.Bind(&p); err != nil {
	}
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

	////////////////////////////
	// buf := bytes.NewBuffer(body)
	// html := buf.String()

	var statuses Response

	if err := json.Unmarshal(body, &statuses); err != nil {
		fmt.Println("aaaa")
		panic(err)
	}
	// decoded, err := json.Marshal(html)
	// decodedAgain, err := json.Unmarshal(string(decoded))

	// s, err := json.Marshal(html)
	// ss := bytes.NewBuffer(s)
	// sss := ss.String()
	fmt.Println(p)
	return c.JSON(http.StatusOK, statuses)
	////////////////////////////
	// returnValue, err := json.Marshal(body)

	// s := []byte(html)
	// if err := json.Unmarshal(s, &p); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+\n", p)
}
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
