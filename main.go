package main

/*
Making a new type "Config" and making it a struct
the struct holds 3 values, all of which are drawn from a json file
*/

type Config struct {
	IPAdress     string `json:"IPAdress"`
	SerialNumber string `json:"SerialNumber"`
	AccessCode   string `json:"AccessCode"`
}

/*

*/
func main(){

err := printer.Connect()
if err != nil {
    panic(err)
}

pool := bambulabs_api.NewPrinterPool()

for i := 0; i < len(Config[]); i++{
	pool.AddPrinter(Config)
}
}



/*"IPAddress": " ",
"SerialNumber": " ",
"AccessCode": " "
}*/
