package main

import (
	"fmt"
)

//Gamer is a struct
type Gamer struct{
	Name string
	Games []string
}

//AddGame is a method
func (gamer *Gamer) AddGame(games string)  {
	gamer.Games = append(gamer.Games,games)
}

func main() {

	gamer := Gamer{Name: "andri"}

	gamer.AddGame("Mario bros")
	gamer.AddGame("Fifa 2020")
	gamer.AddGame("Soccer 2020")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}

}
