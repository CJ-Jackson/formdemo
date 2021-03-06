package forms

import (
	"bytes"
	"github.com/cjtoolkit/form"
	"github.com/cjtoolkit/form/fields"
	html "html/template"
	"regexp"
	"time"
)

const (
	DEMO_FORM_LETTER_ONLY = "demo_form_letter_only"
)

var (
	letterOnly = regexp.MustCompile(`^[a-zA-Z]+$`)
)

func init() {
	form.AddToEnglishLanguageMap(DEMO_FORM_LETTER_ONLY,
		form.BuildLanguageTemplate(`'{{.Label}}' must only be letters.`))
}

type DemoForm struct {
	tpl    *html.Template
	fields []form.FormFieldInterface

	Checked bool

	IdNorm  string
	IdModel int64
	IdErr   error

	TitleNorm  string
	TitleModel string
	TitleErr   error

	ForenameNorm  string
	ForenameModel string
	ForenameErr   error

	SurnameNorm  string
	SurnameModel string
	SurnameErr   error

	TimeZoneNorm  string
	TimeZoneModel string
	TimeZoneErr   error

	TimeNorm  string
	TimeModel time.Time
	TimeErr   error
}

func NewDemoForm() *DemoForm {
	return (&DemoForm{tpl: demoFormHtml, IdModel: 42, TimeModel: time.Now()}).InitField()
}

func (df *DemoForm) InitField() *DemoForm {
	df.fields = []form.FormFieldInterface{
		fields.NewInt("id", "ID", &df.IdNorm, &df.IdModel, &df.IdErr),
		fields.NewString("title", "Title", &df.TitleNorm, &df.TitleModel, &df.TitleErr,
			fields.StringSuffix(&df.IdNorm), fields.StringRequired(""),
			fields.StringInList("", "Mr", "Mrs")),
		fields.NewString("forename", "Forename", &df.ForenameNorm, &df.ForenameModel, &df.ForenameErr,
			fields.StringSuffix(&df.IdNorm), fields.StringMinRune(3, ""),
			fields.StringPattern(letterOnly, DEMO_FORM_LETTER_ONLY)),
		fields.NewString("surname", "Surname", &df.SurnameNorm, &df.SurnameModel, &df.SurnameErr,
			fields.StringSuffix(&df.IdNorm), fields.StringMinRune(3, ""),
			fields.StringPattern(letterOnly, DEMO_FORM_LETTER_ONLY)),
		fields.NewString("timezone", "Timezone", &df.TimeZoneNorm, &df.TimeZoneModel, &df.TimeErr,
			fields.StringSuffix(&df.IdNorm), fields.StringRequired(""),
			fields.StringInList("", "Europe/Jersey", "Europe/Paris", "America/New_York")),
		fields.NewTime("time", "Date and Time", &df.TimeNorm, &df.TimeModel, &df.TimeErr, time.Local,
			fields.DateTimeLocalFormats(), fields.TimeSuffix(&df.IdNorm), fields.TimeUserLocation(&df.TimeZoneModel)),
	}
	return df
}

func (df *DemoForm) Fields() []form.FormFieldInterface {
	return df.fields
}

func (df *DemoForm) IdField() fields.Int {
	return df.fields[0].(fields.Int)
}

func (df *DemoForm) MatchTitle(title string) bool {
	return title == df.TitleNorm
}

func (df *DemoForm) TitleField() fields.String {
	return df.fields[1].(fields.String)
}

func (df *DemoForm) ForenameField() fields.String {
	return df.fields[2].(fields.String)
}

func (df *DemoForm) SurnameField() fields.String {
	return df.fields[3].(fields.String)
}

func (df *DemoForm) TimeZoneField() fields.String {
	return df.fields[4].(fields.String)
}

func (df *DemoForm) MatchTimeZone(timezone string) bool {
	return timezone == df.TimeZoneNorm
}

func (df *DemoForm) TimeField() fields.Time {
	return df.fields[5].(fields.Time)
}

func (df *DemoForm) HTML() html.HTML {
	buf := &bytes.Buffer{}
	defer buf.Reset()
	df.tpl.Execute(buf, df)
	return html.HTML(buf.Bytes())
}
