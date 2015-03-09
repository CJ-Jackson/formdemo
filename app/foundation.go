package app

import (
	"fmt"
	"net/http"

	"github.com/cjtoolkit/dir"
	"github.com/cjtoolkit/form"
	"github.com/cjtoolkit/form/secondlayer/foundation"
)

func foundationFns(w http.ResponseWriter, r *http.Request) {
	f := form.New(foundation.SecondLayer(), "en-GB")

	value := foundationForm{}
	flash := ""

	get := func() {
		fmt.Fprint(w, `<!DOCTYPE html>
<html>
<head>
<title>Foundation Form Demo</title>
<link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/foundation/5.4.7/css/foundation.min.css">
<script src="//cdnjs.cloudflare.com/ajax/libs/foundation/5.4.7/js/vendor/modernizr.js"></script>
<meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
<div class="row"><div class="large-12 columns">
<h1>Form in Foundation</h1><form method="post" novalidate><div class="flash">`, flash, `</div>`)

		f.Render(&value, w)

		fmt.Fprint(w, `<input type="submit" value="Send" class="form-control"></form>
<a href="/">Go back!</a></div></div>
<script src="https://code.jquery.com/jquery-2.1.1.min.js"></script>
<script src="https://code.jquery.com/ui/1.11.2/jquery-ui.js"></script>
<script type="text/javascript">
</script>
</body>
</html>`)
	}

	switch r.Method {
	case "GET", "HEAD":
		get()
	case "POST": // For those without Javascript.
		r.ParseForm()
		if f.MustValidate(r, &value) {
			flash = fmt.Sprintf(`<div class="alert-box success">Hello, %s %s %s!</div>`,
				es(value.Title), es(value.Firstname), es(value.Secondname))
		} else {
			flash = fmt.Sprint(`<div class="alert-box alert">Hi, There a problem with one or more of the fields!</div>`)
		}
		get()
	default:
		dir.ExecDefaultFailFn(r)
	}
}
