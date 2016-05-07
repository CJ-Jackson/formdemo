package http

import (
	"github.com/cjtoolkit/cli"
	"github.com/cjtoolkit/cli/options"
	"fmt"
	"log"
	"net/http"
)

type httpCommand struct {
	address  string
	httpBoot httpBoot
}

func (hc *httpCommand) CommandConfigure(c *cli.Command) {
	c.
		SetName("http:start").SetDescription("Start HTTP Server").
		AddOption("address", "Set server address (address:port)", options.NewString(&hc.address))
}

func (hc *httpCommand) CommandExecute() {
	fmt.Printf("Start HTTP server at '%s'", hc.address)

	log.Panic(http.ListenAndServe(hc.address, hc.httpBoot))
}

func init() {
	cli.RegisterCommand(&httpCommand{address: ":8080"})
}