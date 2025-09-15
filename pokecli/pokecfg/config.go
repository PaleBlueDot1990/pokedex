package pokecfg

const LocationAreaBaseURL string = "https://pokeapi.co/api/v2/location-area/"

type Config struct {
	NextURL string `json:"next,omitempty"`
	PreviousURL string `json:"previous,omitempty"`
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

func InitConfig() *Config {
	return &Config {
		NextURL: LocationAreaBaseURL,
		PreviousURL: "",
	}
}