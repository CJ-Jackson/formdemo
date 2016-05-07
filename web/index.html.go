package web

import html "html/template"

var indexHtml = html.Must(html.New("index").Parse(`
<h1>Form Demo</h1>

<form method="post" novalidate>
	{{ .FormFields.HTML }}
	<input class="form-control" type="submit" value="Submit">
</form>
`))
