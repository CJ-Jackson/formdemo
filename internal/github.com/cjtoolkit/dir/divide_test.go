package dir

import (
	"testing"
)

func TestDevide(t *testing.T) {
	var signal int

	hello := func(remainingPath string) {
		signal = 1
	}

	world := func(remainingPath string) {
		signal = 2
	}

	misc := func(remainingPath string) {
		firstDir, _ := Divide(remainingPath)

		switch firstDir {
		case "/":
			signal = 3
		case "/test":
			signal = 4
		default:
			signal = 5
		}
	}

	route := func(path string) {
		firstDir, remainingPath := Divide(path)

		switch firstDir {
		case "/":
			hello(remainingPath)
		case "/world":
			world(remainingPath)
		case "/misc":
			misc(remainingPath)
		default:
			signal = 6
		}
	}

	route("/")

	if signal != 1 {
		t.Fail()
	}

	route("/world")

	if signal != 2 {
		t.Fail()
	}

	route("/misc")

	if signal != 3 {
		t.Fail()
	}

	route("/misc/test")

	if signal != 4 {
		t.Fail()
	}

	route("/misc/testA")

	if signal != 5 {
		t.Fail()
	}

	route("/testA")

	if signal != 6 {
		t.Fail()
	}
}
