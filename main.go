package main

import (
	"encoding/json"
	"os"

	"github.com/torbenconto/bambulabs_api"
	"github.com/torbenconto/bambulabs_api/light"
)

// Config holds the configuration for the printers
// It is a slice of structs, each containing the host, access code, and serial number
type Config []struct {
	Host       string `json:"host"`
	AccessCode string `json:"access_code"`
	Serial     string `json:"serial_number"`
}

func main() {
	// Open and read json file

	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	//New variable with type Config
	var configs Config

	// Decode the json file into the configs variable
	err = json.NewDecoder(file).Decode(&configs)
	if err != nil {
		panic(err)
	}

	//Create a new pool of printers in the json
	pool := bambulabs_api.NewPrinterPool()

	//iterate over the json file and add all printers to the pool
	for _, printer := range configs {
		pool.AddPrinter(
			&bambulabs_api.PrinterConfig{
				Host:         printer.Host,
				AccessCode:   printer.AccessCode,
				SerialNumber: printer.Serial,
			},
		)
	}

	// Connect to all printers in the pool
	err = pool.ConnectAll()
	if err != nil {
		panic(err)
	}

	// Execute a function across all connected printers in the pool
	// The function turns on the chamber light for each printer
	pool.ExecuteAll(func(p *bambulabs_api.Printer) error {
		return p.LightOn(light.ChamberLight)
	})

	// Disconnect from all printers in the pool to insure no memory leaks
	err = pool.DisconnectAll()
	if err != nil {
		panic(err)
	}

}
