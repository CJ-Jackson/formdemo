package forms

import (
	html "html/template"
)

func ErrorToHtml(err error) html.HTML {
	if nil == err {
		return html.HTML("")
	}

	return html.HTML(`<span class="help-block">` + html.HTMLEscapeString(err.Error()) + `</span>`)
}

func FormGroup(err error, checked bool) html.HTMLAttr {
	class := "form-group"
	if nil != err {
		class += " has-error"
	} else if checked {
		class += " has-success"
	}
	return html.HTMLAttr(class)
}
