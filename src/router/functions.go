package router

import (
	"github.com/cjtoolkit/groot"
	"net/http"
	"github.com/CJ-Jackson/formdemo/src/http_error"
)

func init() {
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http_error.GetFormDemoError().RaiseNotFound()
	})

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http_error.GetFormDemoError().RaiseMethodNotAllowed()
	})
}

func FormDemoRegisterAction(action groot.ActionInterface) {
	router.RegisterAction(action)
}

func FormDemoRouterServe(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
