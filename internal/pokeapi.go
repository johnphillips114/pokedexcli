package pokeapi 

import (
	
)

type Areas struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type Generation struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type GameIndices struct {
	Generation Generation `json:"generation"`
	GameIndex uint `json:"game_index"`
}

type Language struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type Names struct {
	Name string `json:"name"`
	Language Language `json:"language"`
}

type Region struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type Location struct {
	Names []Names `json:"names"`
	Areas []Areas `json:"areas"`
	Region Region `json:"region"`
	Generation Generation `json:"generation"`
	Name string `json:"name"`
	ID uint `json:"id"`
}

type EncounterMethod struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type Version struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type VersionDetails struct {
	Rate uint `json:"rate"`
	MaxChance uint `json:"max_chance"`
	Version Version `json:"version"`
}

type ConditionValues struct {

}

type Method struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type EncounterDetails struct {
	MinLevel uint `json:"min_level"`
	MaxLevel uint `json:"max_level"`
	ConditionValues []ConditionValues `json:"condition_values"`
	Chance uint `json:"chance"`
	Method Method `json:"method"`
}

type EncounterMethodRates struct {
	EncounterMethod []EncounterMethod `json:"encounter_method"`
	VersionDetails []VersionDetails `json:"version_details"`
}

type PokemonEncounters struct {
	Pokemon []Pokemon `json:"pokemon"`
	VersionDetails []VersionDetails `json:"version_details"`
}

type LocationAreas struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	GameIndex uint `json:"game_index"`
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	Location []Location `json:"location"`
	Names []Names `json:"names"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}