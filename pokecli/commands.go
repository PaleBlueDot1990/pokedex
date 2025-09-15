package pokecli

import (
	"github.com/PaleBlueDot1990/pokedex/pokecli/pokecache"
	pokecfg "github.com/PaleBlueDot1990/pokedex/pokecli/pokecfg"
)

type CliCommand struct {
	name        string 
	description string 
	Callback    func(*pokecfg.Config, *pokecache.Cache) error 
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
}

