package pokecli

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	pokecache "github.com/PaleBlueDot1990/pokedex/pokecli/pokecache"
	pokecfg "github.com/PaleBlueDot1990/pokedex/pokecli/pokecfg"
)

func CommandExit(
	cfg *pokecfg.Config, 
	cache *pokecache.Cache,
	pokemons map[string]pokecfg.Pokemon, 
	args []string,
) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(
	cfg *pokecfg.Config, 
	cache *pokecache.Cache, 
	pokemons map[string]pokecfg.Pokemon, 
	args []string,
) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	
	for _, command := range CliCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil 
}

func CommandMapNext(
	cfg *pokecfg.Config, 
	cache *pokecache.Cache, 
	pokemons map[string]pokecfg.Pokemon, 
	args []string,
) error {
	return CommandMap(cfg, cache, pokemons, cfg.NextURL)
}

func CommandMapBack(
	cfg *pokecfg.Config, 
	cache *pokecache.Cache, 
	pokemons map[string]pokecfg.Pokemon, 
	args []string,
) error {
	return CommandMap(cfg, cache, pokemons, cfg.PreviousURL)
}

func CommandMap(
	cfg *pokecfg.Config, 
	cache *pokecache.Cache,
	pokemons map[string]pokecfg.Pokemon, 
	url string,
) error {
	cache.Mu.Lock()
	defer cache.Mu.Unlock()

	var val []byte 
	entry, ok := cache.Entry[url]

	if !ok {
		fmt.Printf("Results not in cache. Making a http call.\n")
		resp, err := http.Get(url)
		if err != nil {
			return err 
		}

		val, err = io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return err 
		}

		cache.Entry[url] = pokecache.CacheEntry{
			CreatedAt: time.Now(),
			Val:       val,
		}
	} else {
		fmt.Printf("Returning cached results.\n")
		val = entry.Val
	}
	
	locations := &pokecfg.Location{}
	err := json.Unmarshal(val, locations)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Printf("%s\n", location.Name)
	}
	
	cfg.NextURL = locations.Next
	cfg.PreviousURL = locations.Previous
	return nil 
}

func CommandExplore(
	cfg *pokecfg.Config, 
	cache *pokecache.Cache, 
	pokemons map[string]pokecfg.Pokemon,
	args []string,
) error {
	cache.Mu.Lock()
	defer cache.Mu.Unlock()

	locationName := args[0]
	fmt.Printf("Exploring %s...\n", locationName)
	url := pokecfg.LocationAreaBaseURL + locationName + "/"

	var val []byte 
	entry, ok := cache.Entry[url]

	if !ok {
		fmt.Printf("Results not in cache. Making a http call.\n")
		resp, err := http.Get(url)
		if err != nil {
			return err 
		}

		val, err = io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return err 
		}

		cache.Entry[url] = pokecache.CacheEntry{
			CreatedAt: time.Now(),
			Val:       val,
		}
	} else {
		fmt.Printf("Returning cached results.\n")
		val = entry.Val
	}
	
	encounters := &pokecfg.Encounter{}
	err := json.Unmarshal(val, encounters)
	if err != nil {
		return err
	}

	fmt.Printf("Found Pokemon:\n")
	for _, encounter := range encounters.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil 
}

func CommandCatch(
	cfg *pokecfg.Config, 
	cache *pokecache.Cache, 
	pokemons map[string]pokecfg.Pokemon,
	args []string,
) error {
	cache.Mu.Lock()
	defer cache.Mu.Unlock()

	pokemonName := args[0]
	_, ok := pokemons[pokemonName]
	if ok {
		fmt.Printf("%s is already caught!\n", pokemonName)
		return nil 
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	url := pokecfg.PokemonBaseURL + pokemonName + "/"

	resp, err := http.Get(url)
	if err != nil {
		return err 
	}

	val, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err 
	}

	pokemon := &pokecfg.Pokemon{}
	err = json.Unmarshal(val, pokemon)
	if err != nil {
		return err
	}

	baseExperience := pokemon.BaseExperience % 10 
	threshold := rand.Intn(10)
	if baseExperience >= threshold {
		pokemons[pokemonName] = *pokemon
		fmt.Printf("%s was caught!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil 
}