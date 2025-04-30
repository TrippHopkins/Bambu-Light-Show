package main

import (
	"github.com/TrippHopkins/Bambu-Light-Show/functions"
	"github.com/torbenconto/bambulabs_api"
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

	configs, err := open()
	if err != nil {
		panic(err)
	}

	//Create a new pool of printers in the json
	pool := bambulabs_api.NewPrinterPool()

	//iterate over the json file and add all printers to the pool
	//In doing so, we set each trait of each item in the json to the corresponding item in the O.G. API
	//This way we can use the API functions on the printers in the pool without having to worry about the json file
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

	functions.Function1(pool)

	functions.Function2(pool)

	functions.Function3(pool)

	// Disconnect from all printers in the pool
	// This is important to do to avoid memory leaks and other issues
	err = pool.DisconnectAll()
	if err != nil {
		panic(err)
	}
}
