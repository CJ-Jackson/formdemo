package forms

import (
	"github.com/cjtoolkit/form"
	html "html/template"
)

var FuncsMap = html.FuncMap{
	"ErrToHtml": ErrorToHtml,
	"IsErr":     IsErr,
	"StrMap":    form.StringSliceToMap,
	"IntMap":    form.IntSliceToMap,
}
