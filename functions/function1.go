package functions

import (
	"fmt"
	"time"

	"github.com/torbenconto/bambulabs_api"
	"github.com/torbenconto/bambulabs_api/light"
)

// flashes lights and moves printers for funzies (Sequentially)
func Function1(pool *bambulabs_api.PrinterPool) error {
	for i := 0; i < 2; i++ {
		//while the i value is an odd number
		if i%2 == 0 {

			//connects to the printer pool and executes the function in the order of the printers in a json file
			pool.ExecuteAllSequentially(func(p *bambulabs_api.Printer) error {

				//sleep for effect
				time.Sleep(1000 * time.Millisecond)

				//turn on the chamber lights
				err := p.LightOn(light.ChamberLight)
				if err != nil {
					fmt.Println("There was an Error With the Lights1")
					panic(err)
				}
				//move the printer bed to the z access inputted
				err = p.SendGcode([]string{"G1 Z200"})
				if err != nil {
					fmt.Println("There was an Error With the GCode1")
					panic(err)
				}

				return nil

			})
		}
		//while the value is an even number
		if i%2 == 1 {
			//connects to the printer pool and executes the function in the order of the printers in a json file
			pool.ExecuteAllSequentially(func(p *bambulabs_api.Printer) error {

				//sleep for effect
				time.Sleep(1000 * time.Millisecond)

				//flash the chamber lights
				err := p.LightFlashing(light.ChamberLight, 100, 100, 10, 100)
				if err != nil {
					fmt.Println("There was an Error With the Lights2")
					panic(err)
				}

				//move the printer bed to the z access inputted
				err = p.SendGcode([]string{"G1 Z15"})
				if err != nil {
					fmt.Println("There was an Error With the Gcode2")
					panic(err)
				}

				return nil
			})
		}
		//sleep for effect
		time.Sleep(10000 * time.Millisecond)
	}
	return nil
}
