package common

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"strings"
	"fmt"
)

type readerWithAnxiety struct{}

func (_ readerWithAnxiety) Read(p []byte) (n int, err error) {
	err = fmt.Errorf("I've anxiety")
	return
}

func TestFunctions(t *testing.T) {

	Convey("ReaderToHtml", t, func() {
		var i int

		_, i = renderToHtml(nil)
		So(i, ShouldEqual, 1)

		_, i = renderToHtml(readerWithAnxiety{})
		So(i, ShouldEqual, 2)

		_, i = renderToHtml(strings.NewReader("Hi"))
		So(i, ShouldEqual, 0)
	})

}