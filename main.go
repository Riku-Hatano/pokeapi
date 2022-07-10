package main

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

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

func showDatumByName(c echo.Context) error {
	name := c.FormValue("name")
	for i := 0; i <= 1140; i++ {
		if name == pokemons[i].Name {
			return c.JSON(http.StatusOK, pokemons[i].Id)
		}
	}
	return c.JSON(http.StatusOK, "missing name")
}

var img []byte
var img2 string

func gopherPNG() io.Reader {
	return base64.NewDecoder(base64.StdEncoding, strings.NewReader(string(img)))
}

func showDatumById(c echo.Context) error {
	name, _ := strconv.Atoi(c.FormValue("number"))
	for i := 0; i < 1140; i++ {
		if name == pokemons[i].Id {
			var returns []string
			returns = append(returns, pokemons[i].Name, pokemons[i].Url)
			// 			////////////////////
			// 			//以下画像を表示させるための処理
			// 			////////////////////////
			// j, err := http.Get("https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/versions/generation-viii/icons/800.png")
			// if err != nil {
			// 	panic(err)
			// }
			// img, err := ioutil.ReadAll(j.Body)
			// if err != nil {
			// 	panic(err)
			// }
			// file, err := os.Create("sample.png")
			// if err != nil {
			// 	panic(err)
			// }
			// defer file.Close()
			// file.Write(img)

			// p, err := os.Open("./sample.png")
			// if err != nil {
			// 	panic(err)
			// }
			// fmt.Println(p)
			// defer p.Close()
			// imgg, err := png.Decode(p)
			// if err != nil {
			// 	panic(err)
			// }
			// fmt.Println(imgg)

			// defer j.Body.Close()

			////////////////////////////////
			//画像を表示させる処理終わり
			/////////////////////////////////
			return c.JSON(http.StatusOK, returns)
		}
	}
	return c.JSON(http.StatusOK, "missing id")
}
