package main

import (
	"encoding/json"
	"os"
	"time"

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

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			pool.ExecuteAllSequentially(func(p *bambulabs_api.Printer) error {

				time.Sleep(1000 * time.Millisecond)

				err := p.LightOn(light.ChamberLight)
				if err != nil {
					return err
				}

				err = p.SendGcode([]string{"G1 Z255"})
				if err != nil {
					return err
				}

				return nil

			})
		}

		if i%2 == 1 {
			pool.ExecuteAllSequentially(func(p *bambulabs_api.Printer) error {
				time.Sleep(1000 * time.Millisecond)

				err := p.LightFlashing(light.ChamberLight, 100, 100, 10, 100)
				if err != nil {
					return err
				}

				err = p.SendGcode([]string{"G1 Z0"})
				if err != nil {
					return err
				}

				return nil
			})
		}
		time.Sleep(10000 * time.Millisecond)

	}

	// pool.ExecuteAllSequentially(func(p *bambulabs_api.Printer) error {
	// 	time.Sleep(1000 * time.Millisecond)
	// 	return p.SendGcode([]string{"G1 Z0"})
	// })

	// pool.ExecuteAll(func(p *bambulabs_api.Printer) error {
	// 	return p.LightFlashing(light.ChamberLight, 50, 50, 10, 50)
	// })

	// for i := 0; i < 1000; i++ {
	// 	pool.ExecuteAll(func(p *bambulabs_api.Printer) error {
	// 		time.Sleep(468 * time.Millisecond)

	// 		if i%2 == 0 {
	// 			return p.LightOn(light.ChamberLight)
	// 		} else {
	// 			return p.LightOff(light.ChamberLight)
	// 		}
	// 	})
	// }

	// Execute a function across all connected printers in the pool
	// The function turns on the chamber light for each printer
	// pool.ExecuteAllSequentially(func(p *bambulabs_api.Printer) error {
	// 	time.Sleep(150 * time.Millisecond)

	// 	return p.LightOff(light.ChamberLight)
	// })

	// Disconnect from all printers in the pool to insure no memory leaks
	err = pool.DisconnectAll()
	if err != nil {
		panic(err)
	}

}
