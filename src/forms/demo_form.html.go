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
		<label for="demo-forename-{{ .IdNorm }}">{{ $forename.Label }}</label>
		<input id="demo-forename-{{ .IdNorm }}" class="form-control" type="text" name="{{ $forename.NameWithSuffix }}" value="{{ .ForenameNorm }}"
			pattern="{{ $forename.Pattern.String }}" >
		{{ ErrToHtml .ForenameErr }}
	</div>

	{{- $surname := .SurnameField -}}
	<div class="col-md-6 form-group {{- if IsErr .SurnameErr }} has-error {{- else if .Checked }} has-success {{- end}}">
		<label for="demo-surname-{{ .IdNorm }}">{{ $surname.Label }}</label>
		<input id="demo-surname-{{ .IdNorm }}" class="form-control" type="text" name="{{ $surname.NameWithSuffix }}" value="{{ .SurnameNorm }}"
			pattern="{{ $surname.Pattern.String }}" >
		{{ ErrToHtml .SurnameErr }}
	</div>
</div>

<div class="row">
	{{- $timezone := .TimeZoneField -}}
	<div class="col-md-6 form-group {{- if IsErr .TitleErr }} has-error {{- else if .Checked }} has-success {{- end}}">
		<label for="demo-timezone-{{ .IdNorm }}">{{ $timezone.Label }}</label>
		<select id="demo-timezone-{{ .IdNorm }}" class="form-control" name="{{ $timezone.NameWithSuffix }}">
			{{ range .TimeZones -}}
			<option value="{{.}}" {{- if $top.MatchTimeZone .}} selected {{- end }}>{{.}}</option>
			{{- end}}
		</select>
		{{ ErrToHtml .TimeZoneErr }}
	</div>

	{{- $time := .TimeField -}}
	<div class="col-md-6 form-group {{- if IsErr .TimeErr }} has-error {{- else if .Checked }} has-success {{- end}}">
		<label for="demo-time-{{ .IdNorm }}">{{ $time.Label }}</label>
		<input id="demo-time-{{ .IdNorm }}" class="form-control" type="time" name="{{ $time.NameWithSuffix }}" value="{{ .TimeNorm }}" >
		{{ ErrToHtml .TimeErr }}
	</div>
</div>
`))
