# CJToolkit Dir

URL/Dir Path Toolkit

Documentation can be found at.

https://godoc.org/github.com/cjtoolkit/dir

## Installation

~~~
go get github.com/cjtoolkit/dir
~~~

## Example

~~~ go
package main

import (
	"fmt"
	"github.com/cjtoolkit/dir"
	"net/http"
)

func router(w http.ResponseWriter, r *http.Request) {
	defer dir.Clear(r)
	e := func() {
		w.WriteHeader(404)
		fmt.Fprint(w, "404: Not Found")
	}

	switch dir.DivideHttpPath(r) {
	case "/":
		dir.B(r, func() {
			world(w, r)
		}, e)
	case "/p":
		dir.L(1, r, func() {
			world(w, r)
		}, e)
	case "/favicon.ico":
		switch r.Method {
		case "GET":
			e()
		default:
			w.WriteHeader(405)
			fmt.Fprint(w, "405: Method not allowed")
		}
	case "/sub":
		anotherRouter(w, r)
	default:
		e()
	}
}


func world(w http.ResponseWriter, r *http.Request) {
	page := int(1)
	dir.Scan(r, &page)

	fmt.Fprint(w, "Hello World Page: ", page)
}

func anotherRouter(w http.ResponseWriter, r *http.Request) {
	e := func() {
		w.WriteHeader(404)
		fmt.Fprint(w, "404: Not Found")
	}

	switch dir.DivideHttpPath(r) {
	case "/hello":
		dir.B(r, func() {
			hello(w, r)
		}, e)
	default:
		e()
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	currentPath, _ := dir.FetchData(r)
	fmt.Fprint(w, "Hello World, Subpage ", currentPath)
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(router))
}

~~~

## Buy me a beer!

Bitcoin - 1MieXR5ANYY6VstNanhuLRtGQGn6zpjxK3