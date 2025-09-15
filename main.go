package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	pokecli "github.com/PaleBlueDot1990/pokedex/pokecli"
	pokecache "github.com/PaleBlueDot1990/pokedex/pokecli/pokecache"
	pokecfg "github.com/PaleBlueDot1990/pokedex/pokecli/pokecfg"
)

func main() {
	pokecli.InitCliCommands()
	cfg := pokecfg.InitConfig()
	cache := pokecache.NewCache(5 * time.Second)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userPrompt := scanner.Text()

		cleanedUserPrompt := pokecli.CleanInput(userPrompt)
		if len(cleanedUserPrompt) == 0 {
			continue 
		}

		commandName := cleanedUserPrompt[0]
		command, ok := pokecli.CliCommands[commandName]
		if !ok {
			fmt.Printf("Unknown command\n")
			continue 
		}
		
		err := command.Callback(cfg, cache)
		if err != nil {
			fmt.Printf("Something wrong happened- %v\n", err)
		}
	}
}
