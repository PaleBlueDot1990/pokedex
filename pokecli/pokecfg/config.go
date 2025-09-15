package pokecfg

const LocationAreaBaseURL string = "https://pokeapi.co/api/v2/location-area/"
const PokemonBaseURL string = "https://pokeapi.co/api/v2/pokemon/"

type Config struct {
	NextURL string 
	PreviousURL string 
}

type Location struct {
	Next string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results []struct {
		Name string `json:"name,omitempty"`
	} `json:"results,omitempty"`
}

type Encounter struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name,omitempty"`
		} `json:"pokemon,omitempty"`
	} `json:"pokemon_encounters,omitempty"`
}

type Pokemon struct {
	Name           string `json:"name,omitempty"`
	BaseExperience int    `json:"base_experience,omitempty"`
	Height         int    `json:"height,omitempty"`
	Weight         int    `json:"weight,omitempty"`

	Abilities      []struct {
		Ability  struct {
			Name string `json:"name,omitempty"`
		} `json:"ability,omitempty"`
	} `json:"abilities,omitempty"`
	
	Stats []struct {
		BaseStat int `json:"base_stat,omitempty"`
		Stat     struct {
			Name string `json:"name,omitempty"`
		} `json:"stat,omitempty"`
	} `json:"stats,omitempty"`

	Types []struct {
		Type struct {
			Name string `json:"name,omitempty"`
		} `json:"type,omitempty"`
	} `json:"types,omitempty"`
}

func InitConfig() *Config {
	return &Config {
		NextURL: LocationAreaBaseURL,
		PreviousURL: "",
	}
}