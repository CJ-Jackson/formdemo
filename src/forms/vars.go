package forms

import (
	"github.com/cjtoolkit/form"
	html "html/template"
)

var FuncsMap = html.FuncMap{
	"ErrToHtml": ErrorToHtml,
	"StrMap":    form.StringSliceToMap,
	"IntMap":    form.IntSliceToMap,
	"FormGroup": FormGroup,
}
