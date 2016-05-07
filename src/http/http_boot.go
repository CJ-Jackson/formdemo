package http

import (
	"net/http"
	"github.com/CJ-Jackson/formdemo/src/http_error"
	"github.com/CJ-Jackson/formdemo/src/router"
)

type httpBoot struct {}

func (b httpBoot) handleError(w http.ResponseWriter, r *http.Request) {
	http_error.NewFormDemoErrorHandler(recover()).GetHttpHandler().ServeHTTP(w, r)
}

func (b httpBoot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer b.handleError(w, r)

	router.FormDemoRouterServe(w, r)
}