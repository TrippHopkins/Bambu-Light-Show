package main

import (
	"os"

	"github.com/torbenconto/bambulabs_api"
	"github.com/torbenconto/bambulabs_api/light"
)

/*
Making a new type "Config" and making it a struct
the struct holds 3 values, all of which are drawn from a json file
*/

type Config struct {
	IPAddress    string `json:"IPAddress"`
	SerialNumber string `json:"SerialNumber"`
	AccessCode   string `json:"AccessCode"`
}

/*
 */
func main() {

	// Open the JSON file
	file, err := os.Open("config.json")

	// Check for errors
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new pool for all printers
	pool := bambulabs_api.NewPrinterPool()

	//new Variable with type Config
	PrinterConfigs := []Config{}

	//Iterate over json file and add all printers to the pool
	//all printers have a IP Address, Serial Number and Access Code
	for _, Config := range PrinterConfigs {
		pool.AddPrinter(&bambulabs_api.PrinterConfig{
			Host:         Config.IPAddress,
			SerialNumber: Config.SerialNumber,
			AccessCode:   Config.AccessCode,
		})
	}

	//
	err = pool.ConnectAll()
	if err != nil {
		panic(err)
	}

	//Telling all the printers in the pool to turn on the chamber light
	//The Fuction "ExecuteAll" requires a return of a printer and reference to the light
	err = pool.ExecuteAll(func(printer *bambulabs_api.Printer) error {
		return printer.LightOn(light.ChamberLight)
	})

	if err != nil {
		panic(err)
	}
	//disconnect from all printers in the pool
	err = pool.DisconnectAll()
	if err != nil {
		panic(err)
	}
}
