package app

import (
	"fmt"
	"net/http"

	"github.com/cjtoolkit/dir"
	_ "github.com/cjtoolkit/form/lang/enGB"
)

func callError(w http.ResponseWriter, r *http.Request) func() {
	return func() {
		switch r.Method {
		case "GET", "HEAD":
			w.WriteHeader(404)
			fmt.Fprint(w, "404: Not Found")
		default:
			fmt.Fprint(w, "405: Method Now Allowed")
		}
	}
}

func router(w http.ResponseWriter, r *http.Request) {
	dir.UseDiscarderOnHead(&w, r)
	dir.SetDefaultFailFn(r, callError(w, r))
	switch dir.DivideHttpPath(r) {
	case "/":
		dir.B(r, func() {
			index(w, r)
		}, nil)

	case "/bootstrap":
		dir.B(r, func() {
			bootstrapFns(w, r)
		}, nil)

	case "/foundation":
		dir.B(r, func() {
			foundationFns(w, r)
		}, nil)

	default:
		dir.ExecDefaultFailFn(r)
	}
	dir.Clear(r)
}

func init() {
	http.DefaultServeMux.HandleFunc("/", router)
}
