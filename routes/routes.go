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

func CORS(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		next(w, r)
	})
}

func handlerGetAllTypes(w http.ResponseWriter, r *http.Request) {
	typespokemon, err := database.GetAllTypes()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(typespokemon)
}

func handlerGetType(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println(err)
	}
	typepokemon, err := database.GetType(id)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(typepokemon)
}

func handlerGetPokemon(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["pokedex"]
	pokemon, err := database.GetPokemon(id)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(pokemon)
}

func handlerGetAllPokemons(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()["page"]) > 0 {
		page, err := strconv.Atoi(r.URL.Query()["page"][0])
		if err != nil {
			fmt.Println(err)
		} else {
			if page >= 1 {
				pokemon, err := database.GetAllPokemons((page - 1) * 20)
				if err != nil {
					fmt.Println(err)
				}
				json.NewEncoder(w).Encode(pokemon)
			} else {
				fmt.Println(err)
			}
		}
	} else {
		pokemon, err := database.GetAllPokemons(0)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(pokemon)
	}
}

func handlerGetAllAbilities(w http.ResponseWriter, r *http.Request) {
	abilitiespokemon, err := database.GetAllAbilities()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(abilitiespokemon)
}

func handlerGetAbility(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println(err)
	}
	abilitypokemon, err := database.GetAbility(id)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(abilitypokemon)
}

func handlerGetPokemonForAbility(w http.ResponseWriter, r *http.Request) {

}

func handlerGetPokemonForType(w http.ResponseWriter, r *http.Request) {

}

func handlerRoutes(w http.ResponseWriter, r *http.Request) {

	routeName := mux.CurrentRoute(r).GetName()

	switch routeName {
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
	}
}

func AppRouter() *mux.Router {

	routes := mux.NewRouter()

	routes.HandleFunc("/api/types", CORS(handlerGetAllTypes)).Methods("GET")
	routes.HandleFunc("/api/types/{id}", CORS(handlerGetType)).Methods("GET")
	routes.HandleFunc("/api/abilities", CORS(handlerGetAllAbilities)).Methods("GET")
	routes.HandleFunc("/api/abilities/{id}", CORS(handlerGetAbility)).Methods("GET")
	routes.HandleFunc("/api/pokemon", CORS(handlerGetAllPokemons)).Methods("GET")
	routes.HandleFunc("/api/pokemon/{pokedex}", CORS(handlerGetPokemon)).Methods("GET")
	//routes.HandleFunc("/api/pokemon/types/{id}", handlerRoutes).Name("GetPokemonsForType").Methods("GET")
	//routes.HandleFunc("/api/pokemon/abilities/{id}", handlerRoutes).Name("GetPokemonsForAbility").Methods("GET")

	return routes
}
