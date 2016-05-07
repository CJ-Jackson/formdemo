package forms

import html "html/template"

var demoFormHtml = html.Must(html.New("form_demo").Funcs(FuncsMap).Parse(`
{{- $top := . -}}

{{- $id := .IdField -}}
<input type="hidden" name="{{ $id.Name }}" value="{{ .IdNorm }}">

{{- $title := .TitleField -}}
<div class="form-group {{- if IsErr .TitleErr }} has-error {{- else if .Checked }} has-success {{- end}}">
	<label for="demo-title">{{ $title.Label }}</label>
	<select id="demo-title" class="form-control" name="{{ $title.NameWithSuffix }}">
		{{ range .Titles -}}
		<option value="{{.}}" {{- if $top.MatchTitle .}} selected {{- end }}>{{.}}</option>
		{{- end}}
	</select>
	{{ ErrToHtml .TitleErr }}
</div>

<div class="row">
	{{- $forename := .ForenameField -}}
	<div class="col-md-6 form-group {{- if IsErr .ForenameErr }} has-error {{- else if .Checked }} has-success {{- end}}">
		<label for="demo-forename">{{ $forename.Label }}</label>
		<input id="demo-forename" class="form-control" type="text" name="{{ $forename.NameWithSuffix }}" value="{{ .ForenameNorm }}"
			pattern="{{ $forename.Pattern.String }}" >
		{{ ErrToHtml .ForenameErr }}
	</div>

	{{- $surname := .SurnameField -}}
	<div class="col-md-6 form-group {{- if IsErr .SurnameErr }} has-error {{- else if .Checked }} has-success {{- end}}">
		<label for="demo-surname">{{ $surname.Label }}</label>
		<input id="demo-surname" class="form-control" type="text" name="{{ $surname.NameWithSuffix }}" value="{{ .SurnameNorm }}"
			pattern="{{ $surname.Pattern.String }}" >
		{{ ErrToHtml .SurnameErr }}
	</div>
</div>

{{- $time := .TimeField -}}
<div class="form-group {{- if IsErr .TimeErr }} has-error {{- else if .Checked }} has-success {{- end}}">
	<label for="demo-time">{{ $time.Label }}</label>
	<input id="demo-time" class="form-control" type="time" name="{{ $time.NameWithSuffix }}" value="{{ .TimeNorm }}" >
	{{ ErrToHtml .TimeErr }}
</div>
`))
