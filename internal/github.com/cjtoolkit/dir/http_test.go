package dir

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	a := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "A")
	}

	al := func(w http.ResponseWriter, r *http.Request) {
		var num1 int64
		var num2 uint64
		var num3 float64

		err := DirScan(r, &num1)

		b, _ := Scan(r, &num2, &num3)
		if !b || err != nil {
			fmt.Fprint(w, "U")
			return
		}

		fmt.Fprint(w, num1, num2, num3)
	}

	b := func(w http.ResponseWriter, r *http.Request) {
		switch DivideHttpPath(r) {
		case "/":
			fmt.Fprint(w, "B")
		}
	}

	main := func(w http.ResponseWriter, r *http.Request) {
		defer Clear(r)
		switch DivideHttpPath(r) {
		case "/":
			a(w, r)
		case "/b":
			b(w, r)
		default:
			L(2, r, func() {
				al(w, r)
			}, nil)
		}

	}

	ts := httptest.NewServer(http.HandlerFunc(main))

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fail()
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fail()
	}

	if string(body) != "A" {
		t.Fail()
	}

	res, err = http.Get(ts.URL + "/jhui")
	if err != nil {
		t.Fail()
	}

	body, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fail()
	}

	if string(body) == "A" {
		t.Fail()
	}

	res, err = http.Get(ts.URL + "/b")
	if err != nil {
		t.Fail()
	}

	body, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fail()
	}

	if string(body) != "B" {
		t.Fail()
	}

	res, err = http.Get(ts.URL + "/20/-20/20")
	if err != nil {
		t.Fail()
	}

	body, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fail()
	}

	if string(body) != "U" {
		t.Fail()
	}

	res, err = http.Get(ts.URL + "/-20/20/20")
	if err != nil {
		t.Fail()
	}

	body, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fail()
	}

	if string(body) != "-20 20 20" {
		t.Fail()
	}

	res, err = http.Get(ts.URL + "/a")
	if err != nil {
		t.Fail()
	}
}
