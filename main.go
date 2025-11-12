package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/torbenconto/bambulabs_api"
	"github.com/torbenconto/bambulabs_api/light"
)

// File: main.go
// Package: main
// Description: This program connects to multiple Bambu Labs printers using their API and turns on the chamber light for each printer.

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
