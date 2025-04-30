package functions

import (
	"time"

	"github.com/torbenconto/bambulabs_api"
	"github.com/torbenconto/bambulabs_api/light"
	"github.com/torbenconto/bambulabs_api/printspeed"
)

// makes the printers form a diagonal line from bottom left to top right and vice versa
func Function2(pool *bambulabs_api.PrinterPool) error {
	for i := 0; i < 4; i++ {
		//Execute the function on the at the index of the printer provided relative to the json file
		pool.ExecuteOnN(func(p *bambulabs_api.Printer) error {

			//set the print speed to sport
			err := p.SetPrintSpeed(printspeed.Sport)
			if err != nil {
				return err
			}
			//turn on the chamber lights
			err = p.LightOff(light.ChamberLight)
			if err != nil {
				return err
			}
			//move the printer bed to the z access inputted
			err = p.SendGcode([]string{"G1 Z100"})
			if err != nil {
				return err
			}

			return nil
			//these are the indexes that are used to execute the function on the printers in the json file
		}, []int{1, 2})

		pool.ExecuteOnN(func(p *bambulabs_api.Printer) error {

			err := p.SetPrintSpeed(printspeed.Sport)
			if err != nil {
				return err
			}

			err = p.LightOn(light.ChamberLight)
			if err != nil {
				return err
			}
			err = p.SendGcode([]string{"G1 Z210"})
			if err != nil {
				return err
			}

			return nil
		}, []int{0, 3})
		time.Sleep(10000 * time.Millisecond)
		pool.ExecuteOnN(func(p *bambulabs_api.Printer) error {
			err := p.SetPrintSpeed(printspeed.Sport)
			if err != nil {
				return err
			}

			err = p.LightOn(light.ChamberLight)
			if err != nil {
				return err
			}
			err = p.SendGcode([]string{"G1 Z210"})
			if err != nil {
				return err
			}

			return nil
		}, []int{1, 2})

		pool.ExecuteOnN(func(p *bambulabs_api.Printer) error {

			err := p.SetPrintSpeed(printspeed.Sport)
			if err != nil {
				return err
			}

			err = p.LightOff(light.ChamberLight)
			if err != nil {
				return err
			}
			err = p.SendGcode([]string{"G1 Z100"})
			if err != nil {
				return err
			}

			return nil
		}, []int{0, 3})
		time.Sleep(15000 * time.Millisecond)
	}
	return nil
}
