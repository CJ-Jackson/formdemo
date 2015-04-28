package dir

import (
	"fmt"
	"net/http"
)

type Decker interface {
	http.Handler
	Push(Decker)
	Call(http.ResponseWriter, *http.Request, string)
	Next(http.ResponseWriter, *http.Request, string)
}

type DeckBase struct {
	next Decker
}

func (d *DeckBase) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer Clear(r)
	d.Call(w, r, DivideHttpPath(r))
}

func (d *DeckBase) Push(deck Decker) {
	if d.next == nil {
		d.next = deck
	} else {
		d.next.Push(deck)
	}
}

func (d *DeckBase) Call(w http.ResponseWriter, r *http.Request, firstDir string) {
	switch firstDir {
	case "/":
		fmt.Fprintln(w, "Hello World.")
	default:
		d.Next(w, r, firstDir)
	}
}

func (d *DeckBase) Next(w http.ResponseWriter, r *http.Request, firstDir string) {
	if d.next != nil {
		d.next.Call(w, r, firstDir)
	} else {
		w.WriteHeader(404)
		fmt.Fprintln(w, "404: Not Found")
	}
}
