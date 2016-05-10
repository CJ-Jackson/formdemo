package forms

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestDemoForm(t *testing.T) {

	form := NewDemoForm()

	form.TitleNorm = "Mr"
	form.TimeZoneNorm = "Europe/Jersey"

	Convey("Check output", t, func() {

		str := string(form.HTML())

		So(strings.Index(str, `value="Mr" selected>`), ShouldNotEqual, -1)
		So(strings.Index(str, `value="Mrs">`), ShouldNotEqual, -1)

		So(strings.Index(str, `value="Europe/Jersey" selected>`), ShouldNotEqual, -1)
		So(strings.Index(str, `value="Europe/Paris">`), ShouldNotEqual, -1)

	})
}
