package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	
	pokemonName := args[0]
	
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("You haven;t caught this pokemon yet")
		
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat) 
	}
	fmt.Printf("Types:\n")
	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s\n", typ.Type.Name) 
	}

	return nil
}