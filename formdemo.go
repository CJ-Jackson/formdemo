package main

import (
	"flag"
	"net/http"
	_ "pj/formdemo/app"
)

var address = flag.String("address", ":8080", "Specify Address")

func main() {
	flag.Parse()
	http.ListenAndServe(*address, nil)
}
