package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	// "pokemonApi2/subpkg"
)

//構造体たち（受け取る側）
type Response struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
type ResponseStats struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
}

//構造体たち（出力する側）
type Pokemons struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Id   int    `json:"id"`
}

type ChosenStats struct {
	BaseStat int `json:"base_stat"`
	Stat     []struct {
		Name string      `json:"name"`
		Url  interface{} `json:"url"`
	} `json:"stat"`
}

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
var statses []ResponseStats

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
	//以下繰り返し処理でポケモンの情報を追加
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
	//繰り返し処理のポケモンの情報の追加終わり
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
func showDatumById(c echo.Context) error {
	name, _ := strconv.Atoi(c.FormValue("number"))
	for i := 0; i < 1140; i++ {
		if name == pokemons[i].Id {
			//////////////////////////////////////
			//種族値、特性表示
			//////////////////////////////////////
			url := "https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(i+1)
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			var stats ResponseStats
			body2, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			if err := json.Unmarshal(body2, &stats); err != nil {
				panic(err)
			}
			////////////////////////////////////
			//種族値表示終わり
			////////////////////////////////////
			////////////////////////////////////
			//pngファイル持ってくる処理
			////////////////////////////////////
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

			////////////////////////////////////
			//pngファイル持ってくる処理終わり
			////////////////////////////////////

			////////////////////////////////////
			//template
			////////////////////////////////////
			var returns []string
			howMany := len(stats.Abilities)
			for i := 0; i < howMany; i++ {
				fmt.Println(i, ": ", stats.Abilities[i].Ability.Name)
			}
			returns = append(returns, pokemons[i].Name, pokemons[i].Url, stats.Stats[0].Stat.Name, strconv.Itoa(stats.Stats[0].BaseStat), stats.Stats[1].Stat.Name, strconv.Itoa(stats.Stats[1].BaseStat), stats.Stats[2].Stat.Name, strconv.Itoa(stats.Stats[2].BaseStat), stats.Stats[3].Stat.Name, strconv.Itoa(stats.Stats[3].BaseStat), stats.Stats[4].Stat.Name, strconv.Itoa(stats.Stats[4].BaseStat), stats.Stats[5].Stat.Name, strconv.Itoa(stats.Stats[5].BaseStat),
				stats.Abilities[0].Ability.Name)
			w := c.Response()
			t, _ := template.ParseFiles("tmpl.html")
			i, _ := strconv.Atoi(returns[3])
			return t.Execute(w, returns)
			/////////////////////////////////////
			//template終わり
			/////////////////////////////////////
		}
	}
	return c.JSON(http.StatusOK, "missing id")
}
