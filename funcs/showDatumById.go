package funcs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pokemonApi2/structs"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func ShowDatumById(c echo.Context) error {
	name, _ := strconv.Atoi(c.FormValue("number"))
	for i := 0; i < 1140; i++ {
		if name == Pokemons[i].Id {
			//////////////////////////////////////
			//種族値、特性表示
			//////////////////////////////////////
			url := "https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(i+1)
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			var stats structs.ResponseStats
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

			var returns []string
			fmt.Println("///////////////////////////////////////////////////////////////////////////////////////////")
			fmt.Println("/////////////////////////////info about searched pokemon in terminal///////////////////////")
			fmt.Println("///////////////////////////////////////////////////////////////////////////////////////////")
			howManyAbilities := len(stats.Abilities)
			for i := 0; i < howManyAbilities; i++ {
				fmt.Println("ability", i, ": ", stats.Abilities[i].Ability.Name)
			}
			fmt.Println("base experience : ", stats.BaseExperience)
			fmt.Println("forms : ", stats.Forms[0].Name)
			if len(stats.GameIndices) != 0 {
				fmt.Println("game indices : ", stats.GameIndices[0].GameIndex, " game version: ", stats.GameIndices[0].Version.Name)
			}
			fmt.Println("height : ", stats.Height)
			fmt.Println("wehght : ", stats.Weight)
			howManyHeldItems := len(stats.HeldItems)
			for i := 0; i < howManyHeldItems; i++ {
				fmt.Println("held item", i, ": ", stats.HeldItems[i].Item.Name)
			}
			fmt.Println("id : ", stats.Id)
			fmt.Println("is default : ", stats.IsDefault)
			fmt.Println("location area encounter : ", stats.LocationAreaEncounters)
			howManyMoves := len(stats.Moves)
			for i := 0; i < howManyMoves; i++ {
				fmt.Println("move : ", stats.Moves[i].Move.Name) //世代によって覚える技が違うので後で調整する(game indiceみたいな感じにやればできるかも)
			}
			fmt.Println("name : ", stats.Name)
			fmt.Println("order : ", stats.Order)
			if len(stats.PastTypes) != 0 {
				fmt.Println("generation : ", stats.PastTypes[0].Generation.Name)
				if len(stats.PastTypes[0].Types) == 1 {
					fmt.Println("past type : ", stats.PastTypes[0].Types[0].Type.Name)
				} else {
					fmt.Println("past type1 : ", stats.PastTypes[0].Types[0].Type.Name)
					fmt.Println("past type2 : ", stats.PastTypes[0].Types[1].Type.Name)
				}
			}
			fmt.Println("species name : ", stats.Species.Name)
			if len(stats.Types) == 1 {
				fmt.Println(stats.Types[0].Type.Name)
			} else {
				fmt.Println(stats.Types[0].Type.Name)
				fmt.Println(stats.Types[1].Type.Name)
			}
			fmt.Println("///////////////////////////////////////////////////////////////////////////////////////////")
			fmt.Println("///////////////////////////////////////////////////////////////////////////////////////////")
			fmt.Println("///////////////////////////////////////////////////////////////////////////////////////////")
			////////////////////////////////////
			//template
			////////////////////////////////////
			returns = append(returns, Pokemons[i].Name, Pokemons[i].Url, stats.Stats[0].Stat.Name, strconv.Itoa(stats.Stats[0].BaseStat), stats.Stats[1].Stat.Name, strconv.Itoa(stats.Stats[1].BaseStat), stats.Stats[2].Stat.Name, strconv.Itoa(stats.Stats[2].BaseStat), stats.Stats[3].Stat.Name, strconv.Itoa(stats.Stats[3].BaseStat), stats.Stats[4].Stat.Name, strconv.Itoa(stats.Stats[4].BaseStat), stats.Stats[5].Stat.Name, strconv.Itoa(stats.Stats[5].BaseStat),
				stats.Abilities[0].Ability.Name)
			w := c.Response()
			t, _ := template.ParseFiles("tmpl.html")
			return t.Execute(w, returns)
			/////////////////////////////////////
			//template終わり
			/////////////////////////////////////
		}
	}
	return c.JSON(http.StatusOK, "missing id")
}