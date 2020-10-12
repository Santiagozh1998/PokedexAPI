package database

import (
	_ "github.com/lib/pq"
)

func getPokemonAbilities(id string) ([]Ability, error) {

	var err error
	var abilities []Ability
	queryAbility :=
		`SELECT Ability.ID_ABILITY, NameAbility
		FROM
		(SELECT *
		FROM AbilityPokemon
		WHERE ID_POKEMON = $1) AS AP INNER JOIN Ability ON Ability.ID_ABILITY = AP.ID_ABILITY;`

	database, err := GetConnection()
	if err != nil {
		return []Ability{}, err
	}

	rows, err := database.Query(queryAbility, id)
	defer rows.Close()

	if err != nil {
		return []Ability{}, err
	} else {

		for rows.Next() {
			var abilitypokemon Ability

			err := rows.Scan(
				&abilitypokemon.IdAbility,
				&abilitypokemon.NameAbility,
			)
			if err != nil {
				return []Ability{}, err
			}

			abilities = append(abilities, abilitypokemon)
		}
	}

	return abilities, nil
}

func getPokemonTypes(id string) ([]Type, error) {

	var err error
	var types []Type
	queryType :=
		`SELECT Type.ID_TYPE, NameType
		FROM
		(SELECT *
		FROM TypePokemon
		WHERE ID_POKEMON = $1) AS TP INNER JOIN Type ON Type.ID_TYPE = TP.ID_TYPE;`

	database, err := GetConnection()
	if err != nil {
		return []Type{}, err
	}

	rows, err := database.Query(queryType, id)
	defer rows.Close()

	if err != nil {
		return []Type{}, err
	} else {

		for rows.Next() {
			var typepokemon Type

			err := rows.Scan(
				&typepokemon.IdType,
				&typepokemon.NameType,
			)
			if err != nil {
				return []Type{}, err
			}

			types = append(types, typepokemon)
		}
	}

	return types, nil
}

func getPokemonWeaknesses(id string) ([]Type, error) {

	var weaknesses []Type
	queryWeakness :=
		`SELECT Type.ID_TYPE, NameType
		FROM
		(SELECT *
		FROM WeaknessPokemon
		WHERE ID_POKEMON = $1) AS WP INNER JOIN Type ON Type.ID_TYPE = WP.ID_TYPE;`

	database, err := GetConnection()
	if err != nil {
		return []Type{}, err
	}

	rows, err := database.Query(queryWeakness, id)
	defer rows.Close()

	if err != nil {
		return []Type{}, err
	} else {

		for rows.Next() {
			var weaknesspokemon Type

			err := rows.Scan(
				&weaknesspokemon.IdType,
				&weaknesspokemon.NameType,
			)
			if err != nil {
				return []Type{}, err
			}

			weaknesses = append(weaknesses, weaknesspokemon)
		}
	}

	return weaknesses, nil
}

func GetAllPokemons() ([]Pokemon, error) {

	var arraypokemon []Pokemon
	var pokemon Pokemon
	database, err := GetConnection()
	if err != nil {
		return []Pokemon{}, err
	}

	queryPokemon :=
		`SELECT *
		FROM Pokemon;`

	rows, err := database.Query(queryPokemon)
	defer rows.Close()

	if err != nil {
		return []Pokemon{}, err
	} else {

		for rows.Next() {

			err := rows.Scan(
				&pokemon.Pokedex,
				&pokemon.Name,
				&pokemon.UrlImage,
				&pokemon.Description,
				&pokemon.Height,
				&pokemon.Weight,
			)
			if err != nil {
				return []Pokemon{}, err
			}

			pokemon.Abilities, err = getPokemonAbilities(pokemon.Pokedex)
			if err != nil {
				return []Pokemon{}, err
			}
			pokemon.Types, err = getPokemonTypes(pokemon.Pokedex)
			if err != nil {
				return []Pokemon{}, err
			}
			pokemon.Weaknesses, err = getPokemonWeaknesses(pokemon.Pokedex)
			if err != nil {
				return []Pokemon{}, err
			}

			arraypokemon = append(arraypokemon, pokemon)
		}
	}

	return arraypokemon, nil
}

func GetPokemon(id string) (Pokemon, error) {

	var pokemon Pokemon
	database, err := GetConnection()
	if err != nil {
		return Pokemon{}, err
	}

	queryPokemon :=
		`SELECT *
		FROM Pokemon
		WHERE ID_POKEMON = $1 OR NamePokemon = $1;`

	err = database.QueryRow(queryPokemon, id).Scan(
		&pokemon.Pokedex,
		&pokemon.Name,
		&pokemon.UrlImage,
		&pokemon.Description,
		&pokemon.Height,
		&pokemon.Weight,
	)
	if err != nil {

		return Pokemon{}, err
	} else {

		pokemon.Abilities, err = getPokemonAbilities(pokemon.Pokedex)
		if err != nil {
			return Pokemon{}, err
		}
		pokemon.Types, err = getPokemonTypes(pokemon.Pokedex)
		if err != nil {
			return Pokemon{}, err
		}
		pokemon.Weaknesses, err = getPokemonWeaknesses(pokemon.Pokedex)
		if err != nil {
			return Pokemon{}, err
		}
	}

	return pokemon, nil
}

func GetPokemonsForType(name int) ([]Pokemon, error) {

	var arraypokemon []Pokemon
	var pokemon Pokemon
	queryType :=
		`SELECT Pokemon.ID_POKEMON, NamePokemon, UrlImage, DescriptionPokemon, HeightPokemon, WeightPokemon
		FROM
		(SELECT ID_POKEMON
		FROM
		(SELECT *
		FROM Type
		WHERE ID_TYPE = $1) AS Temporal
		INNER JOIN TypePokemon ON Temporal.ID_TYPE = TypePokemon.ID_TYPE) AS FilterType
		INNER JOIN Pokemon ON Pokemon.ID_POKEMON = FilterType.ID_POKEMON;`

	database, err := GetConnection()
	if err != nil {
		return []Pokemon{}, err
	}

	rows, err := database.Query(queryType,
		name)
	defer rows.Close()

	if err != nil {
		return []Pokemon{}, err
	} else {

		for rows.Next() {

			err := rows.Scan(
				&pokemon.Pokedex,
				&pokemon.Name,
				&pokemon.UrlImage,
				&pokemon.Description,
				&pokemon.Height,
				&pokemon.Weight,
			)
			if err != nil {
				return []Pokemon{}, err
			} else {

				pokemon.Abilities, err = getPokemonAbilities(pokemon.Pokedex)
				if err != nil {
					return []Pokemon{}, err
				}
				pokemon.Types, err = getPokemonTypes(pokemon.Pokedex)
				if err != nil {
					return []Pokemon{}, err
				}
				pokemon.Weaknesses, err = getPokemonWeaknesses(pokemon.Pokedex)
				if err != nil {
					return []Pokemon{}, err
				}
			}

			arraypokemon = append(arraypokemon, pokemon)
		}
	}

	return []Pokemon{}, err
}

func GetPokemonsForAbility(name int) ([]Pokemon, error) {

	var arraypokemon []Pokemon
	var pokemon Pokemon
	queryAbility :=
		`SELECT Pokemon.ID_POKEMON, NamePokemon, UrlImage, DescriptionPokemon, HeightPokemon, WeightPokemon
		FROM
		(SELECT ID_POKEMON
		FROM
		(SELECT *
		FROM Ability
		WHERE ID_ABILITY = $1) AS Temporal
		INNER JOIN AbilityPokemon ON Temporal.ID_ABILITY = AbilityPokemon.ID_ABILITY) AS FilterAbility
		INNER JOIN Pokemon ON Pokemon.ID_POKEMON = FilterAbility.ID_POKEMON;
		`

	database, err := GetConnection()
	if err != nil {
		return []Pokemon{}, err
	}

	rows, err := database.Query(queryAbility,
		name)
	defer rows.Close()

	if err != nil {
		return []Pokemon{}, err
	} else {

		for rows.Next() {

			err := rows.Scan(
				&pokemon.Pokedex,
				&pokemon.Name,
				&pokemon.UrlImage,
				&pokemon.Description,
				&pokemon.Height,
				&pokemon.Weight,
			)
			if err != nil {
				return []Pokemon{}, err
			} else {

				pokemon.Abilities, err = getPokemonAbilities(pokemon.Pokedex)
				if err != nil {
					return []Pokemon{}, err
				}
				pokemon.Types, err = getPokemonTypes(pokemon.Pokedex)
				if err != nil {
					return []Pokemon{}, err
				}
				pokemon.Weaknesses, err = getPokemonWeaknesses(pokemon.Pokedex)
				if err != nil {
					return []Pokemon{}, err
				}
			}

			arraypokemon = append(arraypokemon, pokemon)
		}
	}

	return arraypokemon, nil
}

func GetAllAbilities() ([]Ability, error) {

	var abilities []Ability

	queryAbility :=
		`SELECT *
		FROM Ability
		ORDER BY ID_ABILITY;`

	database, err := GetConnection()
	if err != nil {
		return []Ability{}, err
	}

	rows, err := database.Query(queryAbility)
	defer rows.Close()

	if err != nil {
		return []Ability{}, err
	} else {

		for rows.Next() {
			var abilityPokemon Ability

			err := rows.Scan(
				&abilityPokemon.IdAbility,
				&abilityPokemon.NameAbility,
			)
			if err != nil {
				return []Ability{}, err
			}

			abilities = append(abilities, abilityPokemon)
		}
	}

	return abilities, nil
}

func GetAllTypes() ([]Type, error) {

	var types []Type
	queryType :=
		`SELECT *
		FROM Type
		ORDER BY ID_TYPE;`

	database, err := GetConnection()
	if err != nil {
		return []Type{}, err
	}

	rows, err := database.Query(queryType)
	defer rows.Close()

	if err != nil {
		return []Type{}, err
	} else {

		for rows.Next() {
			var typePokemon Type

			err := rows.Scan(
				&typePokemon.IdType,
				&typePokemon.NameType,
			)
			if err != nil {
				return []Type{}, err
			}

			types = append(types, typePokemon)
		}
	}

	return types, nil
}

func GetType(id int) (Type, error) {

	var typePokemon Type

	queryType :=
		`SELECT *
		FROM Type
		WHERE ID_TYPE = $1;`

	database, err := GetConnection()
	if err != nil {
		return Type{}, err
	}

	err = database.QueryRow(queryType,
		id).Scan(
		&typePokemon.IdType,
		&typePokemon.NameType,
	)
	if err != nil {
		return Type{}, err
	}

	return typePokemon, nil
}

func GetAbility(id int) (Ability, error) {

	var abilityPokemon Ability

	queryAbility :=
		`SELECT *
		FROM Ability
		WHERE ID_ABILITY = $1;`

	database, err := GetConnection()
	if err != nil {
		return Ability{}, err
	}

	err = database.QueryRow(queryAbility,
		id).Scan(
		&abilityPokemon.IdAbility,
		&abilityPokemon.NameAbility,
	)
	if err != nil {
		return Ability{}, err
	}

	return abilityPokemon, nil
}
