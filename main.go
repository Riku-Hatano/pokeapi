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
var statses []ResponseStats

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
	Name string `json:"name"`
	Url  string `json:"url"`
	Id   int    `json:"id"`
}

// type ResponseStats struct {
// 	Abilities []struct {
// 		Aiblity []struct {
// 			Name string `json:"name"`
// 			Url  string `json:"url"`
// 		} `json:"ability"`
// 		IsHidden bool  `json:"is_hidden"`
// 		Slot     []int `json:"slot"`
// 	} `json:"abilities"`
// 	BaseExperience int `json:"base_experience"`
// 	Forms          []struct {
// 		Name string `json:"name"`
// 		Url  string `json:"url"`
// 	} `json:"forms"`
// 	GameIndices []struct {
// 		GameIndex int `json:"game_index"`
// 		Version   []struct {
// 			Name string `json:"name"`
// 			Url  string `json:"url"`
// 		} `json:"version"`
// 	} `json:"game_index"`
// 	Height    int `json:"height"`
// 	HeldItems []struct {
// 		Item []struct {
// 			Name string `json:"name"`
// 			Url  string `json:"url"`
// 		} `json:"item"`
// 		VersionDetails []*[]struct {
// 			Rarity  int `json:"rarity"`
// 			Version []struct {
// 				Name string `json:"name"`
// 				Url  string `json:"url"`
// 			} `json:"version"`
// 		} `json:"version_details"`
// 	}
// 	Id                     int    `json:"id"`
// 	IsDefault              bool   `json:"is_default"`
// 	LocationAreaEncounters string `json:"location_area_encounters"`
// 	Moves                  []struct {
// 		Move []struct {
// 			Name string `json:"name"`
// 			Url  string `json:"url"`
// 		} `json:"move"`
// 		VersionGroupDetails []*[]struct {
// 			LevelLearnedAt  int `json:"level_learned_at"`
// 			MoveLearnMethod []struct {
// 				Name string `json:"name"`
// 				Url  string `json:"url"`
// 			} `json:"move_learn_method"`
// 			VersionGroup []struct {
// 				Name string `json:"name"`
// 				Url  string `json:"url"`
// 			} `json:"version_group"`
// 		} `json:"version_group_details"`
// 	} `json:"moves"`
// 	Name      string   `json:"name"`
// 	Order     int      `json:"order"`
// 	PastTypes []string `json:"past_types"`
// 	Species   []struct {
// 		Name string `json:"name"`
// 		Url  string `json:"url"`
// 	} `json:"species"`
// 	Spirites interface{}
// Stats    []struct {
// 	BaseStat int `json:"base_stat"`
// 	Effort   int `json:"effort"`
// 	Stat     []struct {
// 		Name string `json:"name"`
// 		Url  string `json:"url"`
// 	} `json:"stat"`
// } `json:"stats"`
// 	Types []struct {
// 		Slot int `json:"slot"`
// 		Type []struct {
// 			Name string `json:"name"`
// 			Url  string `json:"url"`
// 		} `json:"type"`
// 	} `json:"types"`
// 	Weight int `json:"weight"`
// }

type ResponseStats struct {
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
}

type ChosenStats struct {
	BaseStat int `json:"base_stat"`
	Stat     []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"stat"`
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
func showDatumById(c echo.Context) error {
	name, _ := strconv.Atoi(c.FormValue("number"))
	for i := 0; i < 1140; i++ {
		if name == pokemons[i].Id {
			//////////////////////////////////////
			//種族値表示
			////////////////////////////
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
			// for i = 0 ; i < 6 ; i ++ {
			// 	statses = append(statses, ResponseStats{BaseStat: stats[i].})
			// }
			fmt.Println(stats)
			////////////////////////////////////
			//種族値表示終わり
			////////////////////////////////////

			var returns []string
			returns = append(returns, pokemons[i].Name, pokemons[i].Url)

			return c.JSON(http.StatusOK, returns)
			// return c.JSON(http.StatusOK, "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/versions/generation-viii/icons/"+strconv.Itoa(i)+".png")
		}
	}
	return c.JSON(http.StatusOK, "missing id")
}
