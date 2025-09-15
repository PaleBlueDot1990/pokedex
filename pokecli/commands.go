package pokecli

import (
	pokecache "github.com/PaleBlueDot1990/pokedex/pokecli/pokecache"
	pokecfg   "github.com/PaleBlueDot1990/pokedex/pokecli/pokecfg"
)

type CliCommand struct {
	name        string 
	description string 
	Callback    func(
		*pokecfg.Config, 
		*pokecache.Cache, 
		map[string]pokecfg.Pokemon, 
		[]string,
	) error 
}

var CliCommands = map[string]CliCommand{}

func InitCliCommands() {
	CliCommands["exit"] = CliCommand {
		name:        "exit",
		description: "Exit the Pokedex",
		Callback:    CommandExit,
	}

	CliCommands["help"] = CliCommand {
		name:        "help",
		description: "Display help message",
		Callback:    CommandHelp,
	}

	CliCommands["map"] = CliCommand {
		name:        "map",
		description: "Display name of the next 20 location areas",
		Callback:    CommandMapNext,
	}

	CliCommands["mapb"] = CliCommand {
		name:        "mapb",
		description: "Display name of the previous 20 location areas",
		Callback:    CommandMapBack,
	}

	CliCommands["explore"] = CliCommand {
		name:        "explore",
		description: "Display pokemons located in the provided area",
		Callback:    CommandExplore,
	}

	CliCommands["catch"] = CliCommand {
		name:        "catch",
		description: "Catch the provided pokemon",
		Callback:    CommandCatch,
	}

	CliCommands["inspect"] = CliCommand {
		name:        "inspect",
		description: "Print the details of provided pokemon",
		Callback:    CommandInspect,
	}

	CliCommands["pokedex"] = CliCommand {
		name:        "pokedex",
		description: "Print the details of caught pokemons",
		Callback:    CommandPokedex,
	}
}

