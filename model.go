package main

type Employee struct {
	Id         string `form:"id" json:"id"`
	FirstName  string `form:"firstname" json:"firstname"`
	LastName   string `from:"lastname" json:"lastname"`
	Email      string `from:"email" json:"email"`
	Department string `from:"department" json:"department"`
}

type Response struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []Employee `json:"data"`
}

type Pahlawan struct {
	Name          string `json:"name"`
	BirthYear     int    `json:"birth_year"`
	DeathYear     int    `json:"death_year"`
	Description   string `json:"description"`
	AscensionYear int    `json:"ascension_year"`
}

type ResponsePahlawan struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []Pahlawan `json:"list_pahlawan"`
}

// A Response struct to map the Entire Response
type ResponsePokemon struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNumber int            `json:"entry_number"`
	Species     PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ResponseNews struct {
	Link        string `json:"link"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Post        []Post `json:"posts"`
}

type Post struct {
	Link        string `json:"link"`
	Title       string `json:"title"`
	PubDate     string `json:"pubDate"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
}
