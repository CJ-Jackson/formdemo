package web

import (
	"github.com/CJ-Jackson/formdemo/src/forms"
	"github.com/CJ-Jackson/formdemo/src/skeleton/skeleton_mock"
	"github.com/cjtoolkit/form/form_mock"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {

	r := &http.Request{
		Form:     url.Values{"form": {"form"}},
		PostForm: url.Values{"postform": {"postform"}},
	}

	form := form_mock.NewFormMock()
	s := skeleton_mock.NewSkeletonMock()

	demoForm := forms.NewDemoForm()

	testSubject := &IndexAction{
		htmlBody:   indexHtml,
		r:          r,
		form:       form,
		s:          s,
		FormFields: demoForm,
	}

	Convey("Get", t, func(c C) {
		form.C = c
		s.C = c

		go func() {
			form.ExpectTransform(demoForm, true)

			s.ExpectSetTitle("FormDemo 2.0")

			body := string(s.ExpectSetBody())

			c.So(strings.Index(body, `<div class="alert alert-success" role="alert">Perfect, all is well!</div>`),
				ShouldEqual, -1)

			c.So(strings.Index(body, "Name:"), ShouldEqual, -1)
			c.So(strings.Index(body, "Time:"), ShouldEqual, -1)

			c.So(strings.Index(body, `<div class="alert alert-danger" role="alert">Hmm, there a problem with one or two fields!</div>`),
				ShouldEqual, -1)

			s.ExpectExecute()
		}()

		testSubject.Get()
	})

	Convey("Post: Passed Validation", t, func(c C) {
		form.C = c
		s.C = c

		go func() {
			form.ExpectSetForm(r.PostForm)
			form.ExpectValidate(demoForm, true)

			form.ExpectTransform(demoForm, true)

			s.ExpectSetTitle("FormDemo 2.0")

			body := string(s.ExpectSetBody())

			c.So(strings.Index(body, `<div class="alert alert-success" role="alert">Perfect, all is well!</div>`),
				ShouldNotEqual, -1)

			c.So(strings.Index(body, "Name:"), ShouldNotEqual, -1)
			c.So(strings.Index(body, "Time:"), ShouldNotEqual, -1)

			c.So(strings.Index(body, `<div class="alert alert-danger" role="alert">Hmm, there a problem with one or two fields!</div>`),
				ShouldEqual, -1)

			s.ExpectExecute()
		}()

		testSubject.Post()
	})

	testSubject.Success = false
	demoForm.Checked = false

	Convey("Post: Failed Validation", t, func(c C) {
		form.C = c
		s.C = c

		go func() {
			form.ExpectSetForm(r.PostForm)
			form.ExpectValidate(demoForm, false)

			form.ExpectTransform(demoForm, true)

			s.ExpectSetTitle("FormDemo 2.0")

			body := string(s.ExpectSetBody())

			c.So(strings.Index(body, `<div class="alert alert-success" role="alert">Perfect, all is well!</div>`),
				ShouldEqual, -1)

			c.So(strings.Index(body, "Name:"), ShouldEqual, -1)
			c.So(strings.Index(body, "Time:"), ShouldEqual, -1)

			c.So(strings.Index(body, `<div class="alert alert-danger" role="alert">Hmm, there a problem with one or two fields!</div>`),
				ShouldNotEqual, -1)

			s.ExpectExecute()
		}()

		testSubject.Post()
	})

}
