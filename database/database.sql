CREATE TABLE IF NOT EXISTS Pokemon (
	ID_POKEMON VARCHAR(3) primary key,
	NamePokemon text UNIQUE NOT NULL,
	UrlImage text,
	DescriptionPokemon text,
	HeightPokemon text,
	WeightPokemon text
);
CREATE TABLE IF NOT EXISTS Type (
	ID_TYPE serial primary key,
	NameType text NOT NULL
);
CREATE TABLE Ability (
	ID_ABILITY serial primary key,
	NameAbility text NOT NULL
);
CREATE TABLE IF NOT EXISTS TypePokemon (
	ID_POKEMON VARCHAR(3),
	ID_TYPE integer,
	primary key(ID_POKEMON, ID_TYPE),
	foreign key (ID_POKEMON) REFERENCES Pokemon(ID_POKEMON) ON DELETE CASCADE,
	foreign key (ID_TYPE) REFERENCES Type(ID_TYPE) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS WeaknessPokemon (
	ID_POKEMON VARCHAR(3),
	ID_TYPE integer,
	primary key(ID_POKEMON, ID_TYPE),
	foreign key (ID_POKEMON) REFERENCES Pokemon(ID_POKEMON) ON DELETE CASCADE,
	foreign key (ID_TYPE) REFERENCES Type(ID_TYPE) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS AbilityPokemon (
	ID_POKEMON VARCHAR(3),
	ID_ABILITY integer,
	primary key(ID_POKEMON, ID_ABILITY),
	foreign key (ID_POKEMON) REFERENCES Pokemon(ID_POKEMON) ON DELETE CASCADE,
	foreign key (ID_ABILITY) REFERENCES Ability(ID_ABILITY) ON DELETE CASCADE
);