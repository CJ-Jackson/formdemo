package dir

import (
	"io/ioutil"
	"net/http"
)

type discarder struct {
	http.ResponseWriter
	headerCalled bool
}

func (d *discarder) WriteHeader(n int) {
	d.ResponseWriter.WriteHeader(n)
	d.headerCalled = true
}

func (d *discarder) Write(p []byte) (int, error) {
	if !d.headerCalled {
		d.WriteHeader(200)
	}
	return ioutil.Discard.Write(p)
}

// If HTTP Verb is 'HEAD' use change w to use discarder.
func UseDiscarderOnHead(w *http.ResponseWriter, r *http.Request) {
	if r.Method == "HEAD" {
		*w = &discarder{*w, false}
	}
}
