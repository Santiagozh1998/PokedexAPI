package database

type Pokemon struct {
	Pokedex     string
	Name        string
	UrlImage    string
	Description string
	Height      string
	Weight      string
	Types       []Type
	Weaknesses  []Type
	Abilities   []Ability
}

type Type struct {
	IdType   int
	NameType string
}

type Ability struct {
	IdAbility   int
	NameAbility string
}
