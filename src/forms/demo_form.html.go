package forms

import html "html/template"

var demoFormHtml = html.Must(html.New("form_demo").Funcs(FuncsMap).Parse(`
{{- $top := . -}}

{{- $id := .IdField -}}
<input type="hidden" name="{{ $id.Name }}" value="{{ .IdNorm }}">

{{- $title := .TitleField -}}
<div class="{{ FormGroup .TitleErr .Checked }}">
	<label for="demo-title-{{ .IdNorm }}">{{ $title.Label }}</label>
	<select id="demo-title-{{ .IdNorm }}" class="form-control" name="{{ $title.NameWithSuffix }}">
		{{ range $title.InList -}}
		<option value="{{.}}" {{- if $top.MatchTitle .}} selected {{- end }}>{{.}}</option>
		{{- end}}
	</select>
	{{ ErrToHtml .TitleErr }}
</div>

<div class="row">
	{{- $forename := .ForenameField -}}
	<div class="col-md-6 {{ FormGroup .ForenameErr .Checked }}">
		<label for="demo-forename-{{ .IdNorm }}">{{ $forename.Label }}</label>
		<input id="demo-forename-{{ .IdNorm }}" class="form-control" type="text" name="{{ $forename.NameWithSuffix }}" value="{{ .ForenameNorm }}"
			pattern="{{ $forename.Pattern.String }}" >
		{{ ErrToHtml .ForenameErr }}
	</div>

	{{- $surname := .SurnameField -}}
	<div class="col-md-6 {{ FormGroup .SurnameErr .Checked }}">
		<label for="demo-surname-{{ .IdNorm }}">{{ $surname.Label }}</label>
		<input id="demo-surname-{{ .IdNorm }}" class="form-control" type="text" name="{{ $surname.NameWithSuffix }}" value="{{ .SurnameNorm }}"
			pattern="{{ $surname.Pattern.String }}" >
		{{ ErrToHtml .SurnameErr }}
	</div>
</div>

<div class="row">
	{{- $timezone := .TimeZoneField -}}
	<div class="col-md-6 {{ FormGroup .TimeZoneErr .Checked }}">
		<label for="demo-timezone-{{ .IdNorm }}">{{ $timezone.Label }}</label>
		<select id="demo-timezone-{{ .IdNorm }}" class="form-control" name="{{ $timezone.NameWithSuffix }}">
			{{ range $timezone.InList -}}
			<option value="{{.}}" {{- if $top.MatchTimeZone .}} selected {{- end }}>{{.}}</option>
			{{- end}}
		</select>
		{{ ErrToHtml .TimeZoneErr }}
	</div>

	{{- $time := .TimeField -}}
	<div class="col-md-6 {{ FormGroup .TimeErr .Checked }}">
		<label for="demo-time-{{ .IdNorm }}">{{ $time.Label }}</label>
		<input id="demo-time-{{ .IdNorm }}" class="form-control" type="datetime-local" name="{{ $time.NameWithSuffix }}" value="{{ .TimeNorm }}" >
		{{ ErrToHtml .TimeErr }}
	</div>
</div>
`))
