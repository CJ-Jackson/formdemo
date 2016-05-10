package http_error

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestErrorHandler(t *testing.T) {
	Convey("GetHttpHandler", t, func() {
		_, i := NewFormDemoErrorHandler(nil).getHttpHandler()
		So(i, ShouldEqual, 0)

		_, i = NewFormDemoErrorHandler(NewFormDemoErrorHandler(nil)).getHttpHandler()
		So(i, ShouldEqual, 1)
	})
}
