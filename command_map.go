package main

import (
	"fmt"
	"log"

	"github.com/trial-pyth/go-cli/internal/pokeapi"
)

func commandMap() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results{
		fmt.Printf(" - %s", area.Name)
	}

	return nil
}