package routes

import "github.com/Santiagozh1998/PokedexAPI/database"

type PokemonPages struct {
	TotalPages int
	Page       int
	Results    []database.Pokemon
}

type ErrorNotFound struct {
}

type ErrorServerError struct {
}
