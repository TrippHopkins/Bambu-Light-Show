package main

import (
	"encoding/json"
	"os"
)

func open() (Config, error) {

	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	var configs Config

	//New variable with type Config

	// Decode the json file into the configs variable
	err = json.NewDecoder(file).Decode(&configs)
	if err != nil {
		panic(err)
	}
	return configs, nil
}
