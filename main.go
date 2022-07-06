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
	//以下他のapiから情報を持ってきてデータをjson形式に加工（もし情報を取ってくる先のurlの数が増えたら、その分だけ行数が増える）
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

	var response, response2 Response
	var responses []Response

	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body2, &response2); err != nil {
		panic(err)
	}
	responses = append(responses, response, response2)
	var pokemons []Pokemons
	for i := 0; i <= 1; i++ {
		for j := 0; j < 20; j++ {
			pokemons = append(pokemons, Pokemons{Name: responses[i].Results[j].Name, Url: responses[i].Results[j].URL})
			// pokemons = append(pokemons, Response{Results.Resultsname: responses[i].Results[j].Name, URL: responses[i].Results[j].URL})
		}
	}
	// return c.JSON(http.StatusOK, responses)
	// return c.JSON(http.StatusOK, response.Results)
	// return c.JSON(http.StatusOK, responses[0].Results)
	return c.JSON(http.StatusOK, pokemons)
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
type Pokemons struct {
	Name string `json: "name"`
	Url  string `json: "url"`
}

//今できていること。。複数のapiからデータを持ってきてひとつのページに表示
