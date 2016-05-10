package web

import (
	html "html/template"
	"time"
)

var indexHtml = html.Must(html.New("index").Funcs(html.FuncMap{
	"FormatTime": func(t time.Time) string { return t.Format("2006-01-02T15:04 MST") },
}).Parse(`<h1>Form Demo</h1>

{{- if .Success -}}
<div class="alert alert-success" role="alert">Perfect, all is well!</div>
<p>Name: {{ .FormFields.TitleModel }}. {{ .FormFields.ForenameModel }} {{ .FormFields.SurnameModel }}
Time: {{ .FormFields.TimeModel|FormatTime }}</p>
{{- else if .FormFields.Checked -}}
<div class="alert alert-danger" role="alert">Hmm, there a problem with one or two fields!</div>
{{- end -}}

<form method="post" novalidate>
	{{ .FormFields.HTML }}
	<input class="form-control" type="submit" value="Submit">
</form>

<br>
<p><strong>Project URL: </strong> <a href="https://github.com/cjtoolkit/form">https://github.com/cjtoolkit/form</a>
<strong>Source Code of Demo: </strong> <a href="https://github.com/CJ-Jackson/formdemo">https://github.com/CJ-Jackson/formdemo</a></p>
`))
