package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/Santiagozh1998/PokedexAPI/database"
	"github.com/gorilla/mux"
)

func handlerGetAllTypes(w http.ResponseWriter, r *http.Request) {

	var err error
	var typespokemon []database.Type

	typespokemon, err = database.GetAllTypes()
	if err != nil {
		json.NewEncoder(w).Encode(Error{
			Success: false,
			Message: " Internal Server Error. Please try later",
		})
	}
	json.NewEncoder(w).Encode(typespokemon)
}

func handlerGetType(w http.ResponseWriter, r *http.Request) {

	var id int
	var typepokemon database.Type
	var err error

	id, err = strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println(err)
	}
	typepokemon, err = database.GetType(id)
	if err != nil {
		json.NewEncoder(w).Encode(Error{
			Success: false,
			Message: " Internal Server Error. Please try later",
		})
	}
	json.NewEncoder(w).Encode(typepokemon)
}

func handlerGetPokemon(w http.ResponseWriter, r *http.Request) {

	var id string
	var pokemon database.Pokemon
	var err error

	id = mux.Vars(r)["pokedex"]
	pokemon, err = database.GetPokemon(id)
	if err != nil {
		json.NewEncoder(w).Encode(Error{
			Success: false,
			Message: " Internal Server Error. Please try later",
		})
	}
	json.NewEncoder(w).Encode(pokemon)
}

func handlerGetAllPokemons(w http.ResponseWriter, r *http.Request) {

	var page int
	var idtype int
	var idability int
	var err error
	var totalpages int
	var pokemon []database.Pokemon

	page, idtype, idability = GetPropsFromURL(r)

	if idtype != 0 && idability != 0 {

		totalpages, err = database.GetNumberofRowsPokemonforTypeAndAbility(idability, idtype)
		if err != nil {
			fmt.Println(err)
		}
	} else if idtype != 0 {
		totalpages, err = database.GetNumberofRowsPokemonforType(idtype)
		if err != nil {
			fmt.Println(err)
		}
	} else if idability != 0 {
		totalpages, err = database.GetNumberofRowsPokemonforAbility(idability)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		totalpages, err = database.GetNumberofRowsPokemon()
		if err != nil {
			fmt.Println(err)
		}
	}

	totalpages = (totalpages / 20) + 1

	pokemon, err = database.GetAllPokemons((page-1)*20, idtype, idability)
	if err != nil {
		json.NewEncoder(w).Encode(Error{
			Success: false,
			Message: " Internal Server Error. Please try later",
		})
	}

	json.NewEncoder(w).Encode(PokemonPages{
		TotalPages: totalpages,
		Page:       page,
		Results:    pokemon,
	})
}

func handlerGetAllAbilities(w http.ResponseWriter, r *http.Request) {

	var abilitiespokemon []database.Ability
	var err error

	abilitiespokemon, err = database.GetAllAbilities()
	if err != nil {
		json.NewEncoder(w).Encode(Error{
			Success: false,
			Message: " Internal Server Error. Please try later",
		})
	}
	json.NewEncoder(w).Encode(abilitiespokemon)
}

func handlerGetAbility(w http.ResponseWriter, r *http.Request) {

	var id int
	var abilitypokemon database.Ability
	var err error

	id, err = strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println(err)
	}
	abilitypokemon, err = database.GetAbility(id)
	if err != nil {
		json.NewEncoder(w).Encode(Error{
			Success: false,
			Message: " Internal Server Error. Please try later",
		})
	}
	json.NewEncoder(w).Encode(abilitypokemon)
}

func handler404(w http.ResponseWriter, r *http.Request) {

	url := strings.Split(r.URL.String(), "/")
	if url[1] == "api" {

		json.NewEncoder(w).Encode(Error{
			Success: false,
			Message: "The resource you requested could not be found.",
		})
	} else {

		template := template.Must(template.ParseFiles(
			"views/layout.html",
			"views/templates/Error.html"))
		template.ExecuteTemplate(w, "layout", nil)
	}
}

func handlerDocs(w http.ResponseWriter, r *http.Request) {

	template := template.Must(template.ParseFiles(
		"views/layout.html",
		"views/templates/documentation.html"))
	template.ExecuteTemplate(w, "layout", nil)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {

	template := template.Must(template.ParseFiles(
		"views/layout.html",
		"views/templates/index.html"))
	template.ExecuteTemplate(w, "layout", nil)
}

func AppRouter() *mux.Router {

	routes := mux.NewRouter()
	staticFiles := http.FileServer(http.Dir("views/assets/"))
	routes.PathPrefix("/static/").Handler(http.StripPrefix("/static/", staticFiles))

	routes.HandleFunc("/", handlerIndex)
	routes.HandleFunc("/docs", handlerDocs)
	routes.HandleFunc("/api/type", CORS(handlerGetAllTypes)).Methods("GET")
	routes.HandleFunc("/api/type/{id}", CORS(handlerGetType)).Methods("GET")
	routes.HandleFunc("/api/ability", CORS(handlerGetAllAbilities)).Methods("GET")
	routes.HandleFunc("/api/ability/{id}", CORS(handlerGetAbility)).Methods("GET")
	routes.HandleFunc("/api/pokemon", CORS(handlerGetAllPokemons)).Methods("GET")
	routes.HandleFunc("/api/pokemon/{pokedex}", CORS(handlerGetPokemon)).Methods("GET")
	routes.NotFoundHandler = http.HandlerFunc(handler404)

	return routes
}
