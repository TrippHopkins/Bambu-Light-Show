package main

import (
	"os"

	bambulabs_api "github.com/torbenconto/bambulabs_api"
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

	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pool := bambulabs_api.NewPrinterPool()

	configs := []*Config{}
	for _, config := range configs {
		pool.AddPrinter(config)
	}
}

/*"IPAddress": " ",
"SerialNumber": " ",
"AccessCode": " "
}*/
