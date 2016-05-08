package web

import (
	"bytes"
	"fmt"
	"github.com/CJ-Jackson/formdemo/src/forms"
	"github.com/CJ-Jackson/formdemo/src/router"
	"github.com/CJ-Jackson/formdemo/src/skeleton"
	"github.com/cjtoolkit/form"
	html "html/template"
	"net/http"
)

type IndexAction struct {
	tmp        *html.Template
	r          *http.Request
	form       form.FormInterface
	FormFields *forms.DemoForm
	s          skeleton.SkeletonInterface
	flash      string
	Output     string
}

func (i IndexAction) Paths() []string {
	return []string{"/"}
}

func (i IndexAction) New() http.Handler {
	return &IndexAction{}
}

func (i *IndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i.tmp = indexHtml
	i.r = r
	i.form = forms.GetFormDemoForm()
	i.FormFields = forms.NewDemoForm()
	i.s = skeleton.GetFormDemoSkeleton(skeleton.SkeletonResponseWriter(w))
}

func (i *IndexAction) Get() {
	i.form.Transform(i.FormFields)

	body := &bytes.Buffer{}
	i.tmp.Execute(body, i)

	i.s.SetTitle("FormDemo 2.0")
	i.s.SetBody(body)
	i.s.Execute()
}

func (i *IndexAction) Post() {
	i.r.ParseForm()
	i.form.SetForm(i.r.PostForm)
	i.FormFields.Checked = true
	if i.form.Validate(i.FormFields) {
		i.flash = `<div class="alert alert-success" role="alert">Perfect, all is well!</div>`
		i.Output = fmt.Sprintf("Name: %s. %s %s Time: %s",
			i.FormFields.TitleModel, i.FormFields.ForenameModel, i.FormFields.SurnameModel,
			i.FormFields.TimeModel.Format("2006-01-02T15:04 MST"))
	} else {
		i.flash = `<div class="alert alert-danger" role="alert">Hmm, there a problem with one or two fields!</div>`
	}
	i.Get()
}

func (i *IndexAction) Flash() html.HTML {
	return html.HTML(i.flash)
}

func init() {
	router.FormDemoRegisterAction(IndexAction{})
}
