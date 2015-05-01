package app

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	get := func() {
		fmt.Fprint(w, `<!DOCTYPE html>
<html>
<head>
<title>CJToolkit Form Demo</title>
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.0/css/bootstrap.min.css">
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.0/css/bootstrap-theme.min.css">
<meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
<div class="container">
<h1>CJToolkit Form Demo, Select a CSS Framework</h1>
<ul>
<li><a href="/bootstrap">Bootstrap (*)</a></li>
<li><a href="/foundation">Foundation</a></li>
</ul>

(*) Uses AJAX via PUT on Javascript enabled web browser!<br>
<a href="https://github.com/CJ-Jackson/formdemo">https://github.com/CJ-Jackson/formdemo</a>

</div>
<script src="https://code.jquery.com/jquery-2.1.1.min.js"></script>
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.0/js/bootstrap.min.js"></script>
</body>
</html>`)
	}

	switch r.Method {
	case "GET", "HEAD":
		get()
	default:
		callError(w, r)()
	}
}
