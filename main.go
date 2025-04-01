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
	"fmt"
	"os"

	"github.com/torbenconto/bambulabs_api"
	"github.com/torbenconto/bambulabs_api/light"
)

type Config []struct {
	Host       string `json:"host"`
	AccessCode string `json:"access_code"`
	Serial     string `json:"serial_number"`
}

func main() {
	// read json file
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	// read files
	var configs Config

	err = json.NewDecoder(file).Decode(&configs)
	if err != nil {
		panic(err)
	}

	fmt.Println(configs)

	pool := bambulabs_api.NewPrinterPool()

	for _, printer := range configs {
		pool.AddPrinter(
			&bambulabs_api.PrinterConfig{
				Host:         printer.Host,
				AccessCode:   printer.AccessCode,
				SerialNumber: printer.Serial,
			},
		)
	}

	err = pool.ConnectAll()
	if err != nil {
		panic(err)
	}

	// printers := pool.GetPrinters()

	// for _, printer := range printers {
	//     time.Sleep(250 * time.Millisecond)
	//     printer.LightOn(light.ChamberLight)
	// }

	pool.ExecuteAll(func(p *bambulabs_api.Printer) error {
		return p.LightOn(light.ChamberLight)
	})

}
