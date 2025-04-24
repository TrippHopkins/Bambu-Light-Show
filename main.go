package main

import (
	"encoding/json"
	"fmt"
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

	for i := 1; i < 4; i++ {
		if i%2 == 0 {
			pool.ExecuteAllSequentially(func(p *bambulabs_api.Printer) error {

				time.Sleep(1000 * time.Millisecond)

				err := p.LightOn(light.ChamberLight)
				if err != nil {
					fmt.Println("There was an Error With the Lights1")
					panic(err)
				}

				err = p.SendGcode([]string{"G1 Z200"})
				if err != nil {
					fmt.Println("There was an Error With the GCode1")
					panic(err)
				}

				return nil

			})
		}
		if i%2 == 1 {
			pool.ExecuteAllSequentially(func(p *bambulabs_api.Printer) error {
				time.Sleep(1000 * time.Millisecond)

				err := p.LightFlashing(light.ChamberLight, 100, 100, 10, 100)
				if err != nil {
					fmt.Println("There was an Error With the Lights2")
					panic(err)
				}

				err = p.SendGcode([]string{"G1 Z15"})
				if err != nil {
					fmt.Println("There was an Error With the Gcode2")
					panic(err)
				}

				return nil
			})
		}
		time.Sleep(7000 * time.Millisecond)
	}
	// for i := 1; i < 4; i++ {
	// 	if i%2 == 0 {
	// 		pool.ExecuteOnN(func(p *bambulabs_api.Printer) error {
	// 			err = p.LightOn(light.ChamberLight)
	// 			if err != nil {
	// 				return err
	// 			}
	// 			err = p.SendGcode([]string{"G1 Z100"})
	// 			if err != nil {
	// 				return err
	// 			}

	// 			return nil
	// 		}, []int{1, 2})

	// 	}
	// 	pool.ExecuteOnN(func(p *bambulabs_api.Printer) error {
	// 		err = p.LightOff(light.ChamberLight)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		err = p.SendGcode([]string{"G1 Z210"})
	// 		if err != nil {
	// 			return err
	// 		}

	// 		return nil
	// 	}, []int{0, 3})

	// 	if i%2 == 1 {
	// 		pool.ExecuteOnN(func(p *bambulabs_api.Printer) error {
	// 			err = p.LightOff(light.ChamberLight)
	// 			if err != nil {
	// 				return err
	// 			}
	// 			err = p.SendGcode([]string{"G1 Z210"})
	// 			if err != nil {
	// 				return err
	// 			}

	// 			return nil
	// 		}, []int{1, 2})

	// 	}
	// 	pool.ExecuteOnN(func(p *bambulabs_api.Printer) error {
	// 		err = p.LightOn(light.ChamberLight)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		err = p.SendGcode([]string{"G1 Z100"})
	// 		if err != nil {
	// 			return err
	// 		}

	// 		return nil
	// 	}, []int{0, 3})

	// }
	err = pool.DisconnectAll()
	if err != nil {
		panic(err)
	}
}
