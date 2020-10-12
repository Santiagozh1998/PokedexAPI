package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Santiagozh1998/PokedexAPI/database"
	"github.com/gorilla/mux"
)

func handlerRoutes(w http.ResponseWriter, r *http.Request) {

	routeName := mux.CurrentRoute(r).GetName()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	switch routeName {
	case "GetAllPokemons":
		pokemon, err := database.GetAllPokemons()
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(pokemon)
		break
	case "GetPokemon":
		id := string(mux.Vars(r)["pokedex"])
		pokemon, err := database.GetPokemon(id)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(pokemon)
		break
	case "GetPokemonsForType":
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			log.Println(err)
		}
		pokemon, err := database.GetPokemonsForType(id)
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(pokemon)
		break
	case "GetPokemonsForAbility":
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			fmt.Println(err)
		}
		pokemon, err := database.GetPokemonsForAbility(id)
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(pokemon)
		break
	case "GetAllAbilities":
		abilities, err := database.GetAllAbilities()
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(abilities)
		break
	case "GetAbility":
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			fmt.Println(err)
		}
		ability, err := database.GetAbility(id)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(ability)
		break
	case "GetAllTypes":
		types, err := database.GetAllTypes()
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(types)
		break
	case "GetType":
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			fmt.Println(err)
		}
		typepokemon, err := database.GetType(id)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(typepokemon)
		break
	}
}

func AppRouter() *mux.Router {

	routes := mux.NewRouter()

	routes.HandleFunc("/api/types/all", handlerRoutes).Name("GetAllTypes").Methods("GET")
	routes.HandleFunc("/api/types/{id}", handlerRoutes).Name("GetType").Methods("GET")
	routes.HandleFunc("/api/abilities/all", handlerRoutes).Name("GetAllAbilities").Methods("GET")
	routes.HandleFunc("/api/abilities/{id}", handlerRoutes).Name("GetAbility").Methods("GET")
	routes.HandleFunc("/api/pokemon/all", handlerRoutes).Name("GetAllPokemons").Methods("GET")
	routes.HandleFunc("/api/pokemon/{pokedex}", handlerRoutes).Name("GetPokemon").Methods("GET")
	routes.HandleFunc("/api/pokemon/types/{id}", handlerRoutes).Name("GetPokemonsForType").Methods("GET")
	routes.HandleFunc("/api/pokemon/abilities/{id}", handlerRoutes).Name("GetPokemonsForAbility").Methods("GET")

	return routes
}
