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

			url2 := stats.Species.Url
			res2, err := http.Get(url2)
			if err != nil {
				panic(err)
			}
			var stats2 structs.ResponseStats2
			body3, err := ioutil.ReadAll(res2.Body)
			if err != nil {
				panic(err)
			}
			defer res2.Body.Close()
			if err := json.Unmarshal(body3, &stats2); err != nil {
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
			// howManyAbilities := len(stats.Abilities)
			// for i := 0; i < howManyAbilities; i++ {
			// 	fmt.Println("ability", i, ": ", stats.Abilities[i].Ability.Name)
			// }
			// fmt.Println("base experience : ", stats.BaseExperience)
			// fmt.Println("forms : ", stats.Forms[0].Name)
			// if len(stats.GameIndices) != 0 {
			// 	fmt.Println("game indices : ", stats.GameIndices[0].GameIndex, " game version: ", stats.GameIndices[0].Version.Name)
			// }
			// fmt.Println("height : ", stats.Height)
			// fmt.Println("wehght : ", stats.Weight)
			// howManyHeldItems := len(stats.HeldItems)
			// for i := 0; i < howManyHeldItems; i++ {
			// 	fmt.Println("held item", i, ": ", stats.HeldItems[i].Item.Name)
			// }
			// fmt.Println("id : ", stats.Id)
			// fmt.Println("is default : ", stats.IsDefault)
			// fmt.Println("location area encounter : ", stats.LocationAreaEncounters)
			// howManyMoves := len(stats.Moves)
			// for i := 0; i < howManyMoves; i++ {
			// 	fmt.Println("move : ", stats.Moves[i].Move.Name) //世代によって覚える技が違うので後で調整する(game indiceみたいな感じにやればできるかも)
			// }
			// fmt.Println("name : ", stats.Name)
			// fmt.Println("order : ", stats.Order)
			// if len(stats.PastTypes) != 0 {
			// 	fmt.Println("generation : ", stats.PastTypes[0].Generation.Name)
			// 	if len(stats.PastTypes[0].Types) == 1 {
			// 		fmt.Println("past type : ", stats.PastTypes[0].Types[0].Type.Name)
			// 	} else {
			// 		fmt.Println("past type1 : ", stats.PastTypes[0].Types[0].Type.Name)
			// 		fmt.Println("past type2 : ", stats.PastTypes[0].Types[1].Type.Name)
			// 	}
			// }
			// fmt.Println("species name : ", stats.Species.Name)
			// if len(stats.Types) == 1 {
			// 	fmt.Println(stats.Types[0].Type.Name)
			// } else {
			// 	fmt.Println(stats.Types[0].Type.Name)
			// 	fmt.Println(stats.Types[1].Type.Name)
			// }
			fmt.Println("base_happiness : ", stats2.BaseHappiness)
			fmt.Println("capture_rate : ", stats2.CaptureRate)
			fmt.Println("Color : ", stats2.Color.Name)
			howManyEggGroups := len(stats2.EggGroups)
			for i := 0; i < howManyEggGroups; i++ {
				fmt.Println("egg_goups", i+1, " : ", stats2.EggGroups[i].Name)
			}
			fmt.Println("evolution_chain : ", stats2.EvolutionChain.Url)
			fmt.Println("evoluves_form_species : ", stats2.EvoluvesFromSpecies.Name)
			howManyTexts := len(stats2.FlavorTextEntries)
			for i := 0; i < howManyTexts; i++ {
				if stats2.FlavorTextEntries[i].Language.Name == "ja" {
					fmt.Println("flavor_text_entries", i+1, " : ", stats2.FlavorTextEntries[i].FlavorText)
				}
			}
			fmt.Println("form_descriptions : ", stats2.FormDescriptions)
			fmt.Println("forms_sweitchable : ", stats2.FormsSwitchable)
			fmt.Println("gender_rate : ", stats2.GenderRate)
			howManyGeneras := len(stats2.Genera)
			for i := 0; i < howManyGeneras; i++ {
				fmt.Println("genera", i+1, " : ", stats2.Genera[i].Genus)
			}
			fmt.Println("generation : ", stats2.Generation.Name)
			fmt.Println("growth_rate : ", stats2.GrowthRate.Name)
			fmt.Println("habitat : ", stats2.Habitat.Name)
			fmt.Println("has_gender_differences : ", stats2.HasGenderDifferences)
			fmt.Println("hatch_counter : ", stats2.HatchCounter)
			fmt.Println("Id : ", stats2.Id)
			fmt.Println("is_baby : ", stats2.IsBaby)
			fmt.Println("is_legendary : ", stats2.IsLegendary)
			fmt.Println("is_mythical : ", stats2.IsMythical)
			fmt.Println("name : ", stats2.Name)
			howManyNames := len(stats2.Names)
			for i := 0; i < howManyNames; i++ {
				fmt.Println("names", i+1, " : ", stats2.Names[i].Name)
			}
			fmt.Println("order : ", stats2.Order)
			howManyPal := len(stats2.PalParkEncounters)
			for i := 0; i < howManyPal; i++ {
				fmt.Println("pal_park_encounters", i+1, " : ", stats2.PalParkEncounters[i].Area.Name)
			}
			howManyPokedexes := len(stats2.PokedexNumbers)
			for i := 0; i < howManyPokedexes; i++ {
				fmt.Println(stats2.PokedexNumbers[i].Pokedex.Name, " : ", stats2.PokedexNumbers[i].EntryNumber)
			}
			fmt.Println(stats2.Shape.Name)
			howManyVarieties := len(stats2.Varieties)
			for i := 0; i < howManyVarieties; i++ {
				fmt.Println("varieties", i+1, " : ", stats2.Varieties[i].Pokemon.Name)
			}
			fmt.Println("///////////////////////////////////////////////////////////////////////////////////////////")
			fmt.Println("///////////////////////////////////////////////////////////////////////////////////////////")
			fmt.Println("///////////////////////////////////////////////////////////////////////////////////////////")
			////////////////////////////////////
			//template
			////////////////////////////////////
			returns = append(returns, strconv.Itoa(stats.Id), stats.Name, strconv.Itoa(stats.Stats[0].BaseStat), strconv.Itoa(stats.Stats[1].BaseStat), strconv.Itoa(stats.Stats[2].BaseStat), strconv.Itoa(stats.Stats[3].BaseStat), strconv.Itoa(stats.Stats[4].BaseStat), strconv.Itoa(stats.Stats[5].BaseStat),
				url, stats.Species.Url)
			// strconv.Itoa(stats.Stats[1].BaseStat)
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
