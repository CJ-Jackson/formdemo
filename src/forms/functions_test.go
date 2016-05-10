package forms

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFunctions(t *testing.T) {
	Convey("ErrorToHtml", t, func() {
		So(ErrorToHtml(nil), ShouldBeEmpty)
		So(ErrorToHtml(fmt.Errorf("Hi")), ShouldEqual, `<span class="help-block">Hi</span>`)
	})

	Convey("FormGroup", t, func() {
		err := fmt.Errorf("Hi")
		So(FormGroup(nil, false), ShouldEqual, "form-group")
		So(FormGroup(err, false), ShouldEqual, "form-group has-error")
		So(FormGroup(err, true), ShouldEqual, "form-group has-error")
		So(FormGroup(nil, true), ShouldEqual, "form-group has-success")
	})
}
