package test

import (
	"fmt"

	"github.com/Santiagozh1998/PokedexAPI/database"
)

func TestDatabase() {

	var pokemonerror []string
	totalpokemon, err := database.GetNumberofRowsPokemon()
	if err != nil {
		fmt.Println(err)
	}

	for i := 1; i <= totalpokemon; i++ {

		var id string

		if i < 10 {
			id = fmt.Sprintf("00%d", i)
		} else if i < 100 {
			id = fmt.Sprintf("0%d", i)
		} else {
			id = fmt.Sprintf("%d", i)
		}

		pokemon, err := database.GetPokemon(id)
		if err != nil {
			fmt.Println(err)
			pokemonerror = append(pokemonerror, id)
		} else {
			fmt.Println(pokemon.Pokedex, pokemon.Name)
		}
	}

	fmt.Println(pokemonerror)
}
