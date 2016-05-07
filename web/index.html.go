package web

import html "html/template"

var indexHtml = html.Must(html.New("index").Parse(`
<h1>Form Demo</h1>

{{ .Flash }}

<p>{{ .Output }}</p>

<form method="post" novalidate>
	{{ .FormFields.HTML }}
	<input class="form-control" type="submit" value="Submit">
</form>

<br>
<p><strong>Project URL: </strong> <a href="https://github.com/cjtoolkit/form">https://github.com/cjtoolkit/form</a></p>
`))
