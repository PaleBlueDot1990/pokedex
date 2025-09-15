package pokecfg

const LocationAreaBaseURL string = "https://pokeapi.co/api/v2/location-area/"

type Config struct {
	NextURL     string 
	PreviousURL string
}

type Location struct {
	Count    int      `json:"count,omitempty"`
	Next     string   `json:"next,omitempty"`
	Previous string   `json:"previous,omitempty"`
	Results []struct {
		Name string   `json:"name,omitempty"`
		URL  string   `json:"url,omitempty"`
	}
}

func InitConfig() *Config {
	return &Config {
		NextURL: LocationAreaBaseURL,
		PreviousURL: "",
	}
}