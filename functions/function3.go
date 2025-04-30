package functions

import (
	"time"

	"github.com/torbenconto/bambulabs_api"
)

// makes printers alternate between these two patterns
// UP - DOWN - DOWN- UP
// DOWN - UP - UP - DOWN
func Function3(pool *bambulabs_api.PrinterPool) error {

	for i := 0; i < 4; i++ {
		if i%2 == 0 {

			time.Sleep(7500 * time.Millisecond)

			pool.ExecuteOnN(func(p *bambulabs_api.Printer) error {

				p.SendGcode([]string{"G1 Z50"})

				return nil
			}, []int{0})
			pool.ExecuteOnN(func(p *bambulabs_api.Printer) error {

				p.SendGcode([]string{"G1 Z100"})

				return nil
			}, []int{1})
			pool.ExecuteOnN(func(p *bambulabs_api.Printer) error {

				p.SendGcode([]string{"G1 Z150"})

				return nil
			}, []int{2})
			pool.ExecuteOnN(func(p *bambulabs_api.Printer) error {

				p.SendGcode([]string{"G1 Z200"})

				return nil
			}, []int{3})
		}
		if i%2 == 1 {

			time.Sleep(7500 * time.Millisecond)

			pool.ExecuteOnN(func(p *bambulabs_api.Printer) error {

				p.SendGcode([]string{"G1 Z200"})

				return nil
			}, []int{0})
			pool.ExecuteOnN(func(p *bambulabs_api.Printer) error {

				p.SendGcode([]string{"G1 Z150"})

				return nil
			}, []int{1})
			pool.ExecuteOnN(func(p *bambulabs_api.Printer) error {

				p.SendGcode([]string{"G1 Z100"})

				return nil
			}, []int{2})
			pool.ExecuteOnN(func(p *bambulabs_api.Printer) error {

				p.SendGcode([]string{"G1 Z50"})

				return nil
			}, []int{3})
		}
	}

	return nil
}
