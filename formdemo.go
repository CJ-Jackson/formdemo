package main

import (
	"flag"
	_ "github.com/CJ-Jackson/formdemo/app"
	"net/http"
)

var address = flag.String("address", ":8080", "Specify Address")

func main() {
	flag.Parse()
	http.ListenAndServe(*address, nil)
}
