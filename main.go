// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"

// 	"github.com/torbenconto/bambulabs_api"
// 	"github.com/torbenconto/bambulabs_api/light"
// )

// /*
// Making a new type "Config" and making it a struct
// the struct holds 3 values, all of which are drawn from a json file
// */

// type Config []struct {
// 	Host       string `json:"host"`
// 	AccessCode string `json:"access_code"`
// 	Serial     string `json:"serial_number"`
// }

// /*
//  */
// func main() {

// 	// Open the JSON file
// 	file, err := os.Open("config.json")
// 	// Check for errors
// 	if err != nil {
// 		panic(err)
// 	}

// 	var configs Config

// 	// Create a new decoder and decode the JSON into the Config struct

// 	err = json.NewDecoder(file).Decode(&configs)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Create a new pool for all printers
// 	pool := bambulabs_api.NewPrinterPool()

// 	//new Variable with type Config

// 	//Iterate over json file and add all printers to the pool
// 	//all printers have a IP Address, Serial Number and Access Code
// 	for _, printer := range configs {
// 		pool.AddPrinter(&bambulabs_api.PrinterConfig{
// 			Host:         printer.Host,
// 			AccessCode:   printer.AccessCode,
// 			SerialNumber: printer.Serial,
// 		})
// 	}

// 	//
// 	pool.ConnectAll()

// 	//Telling all the printers in the pool to turn on the chamber light
// 	//The Fuction "ExecuteAll" requires a return of a printer and reference to the light
// 	err = pool.ExecuteAll(func(p *bambulabs_api.Printer) error {
// 		return p.LightOn(light.ChamberLight)
// 	})
// 	if err != nil {
// 		fmt.Printf("Error during ExecuteAll: %v\n", err)
// 	}
// 	//disconnect from all printers in the pool
// 	pool.DisconnectAll()
// }

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
