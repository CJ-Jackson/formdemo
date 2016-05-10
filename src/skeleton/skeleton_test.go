package skeleton

import (
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestSkeleton(t *testing.T) {
	Convey("Test output", t, func() {
		testSubject := GetFormDemoSkeleton()

		testSubject.SetTitle("This is the title")
		testSubject.SetBody(strings.NewReader("<p>This is the body</p>"))

		buf := &bytes.Buffer{}

		testSubject.execute(buf)

		str := buf.String()
		buf.Reset()

		So(strings.Index(str, "<title>This is the title"), ShouldNotEqual, -1)
		So(strings.Index(str, "<p>This is the body</p>"), ShouldNotEqual, -1)
	})
}
