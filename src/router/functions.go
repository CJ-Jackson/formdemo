package router

import (
	"github.com/cjtoolkit/groot"
	"net/http"
)

func FormDemoRegisterAction(action groot.ActionInterface) {
	router.RegisterAction(action)
}

func FormDemoRouterServe(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
