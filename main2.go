package main

// import (
// 	"fmt"
// 	"image/png"
// 	"io/ioutil"
// 	"net/http"
// 	"os"

// 	"github.com/labstack/echo/v4"
// )

// func main2(c echo.Context) error {
// 	j, err := http.Get("https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/versions/generation-viii/icons/1.png")
// 	if err != nil {
// 		panic(err)
// 	}
// 	img, _ := ioutil.ReadAll(j.Body)
// 	fmt.Println(img)
// 	file, err := os.Create("sample.png")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	file.Write(img)

// 	raw, err := os.Open("sample.png")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer raw.Close()
// 	wantReturn, _ := png.Decode(raw)
// 	fmt.Println(wantReturn)
// 	return c.JSON(http.StatusOK, wantReturn)
// }
