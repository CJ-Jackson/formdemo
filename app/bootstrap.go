package app

import (
	"fmt"
	"net/http"

	"github.com/cjtoolkit/dir"
	"github.com/cjtoolkit/form"
	"github.com/cjtoolkit/form/secondlayer/bootstrap"
)

func bootstrapFns(w http.ResponseWriter, r *http.Request) {
	f := form.New(bootstrap.SecondLayer(), "en-GB")

	value := bootStrapForm{}
	flash := ""

	get := func() {
		fmt.Fprint(w, `<!DOCTYPE html>
<html>
<head>
<title>Bootstrap Form Demo</title>
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.0/css/bootstrap.min.css">
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.0/css/bootstrap-theme.min.css">
<meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
<div class="container">
<h1>Form in Bootstrap</h1><form method="post" novalidate><div class="flash">`, flash, `</div>`)

		f.Render(&value, w)

		fmt.Fprint(w, `<input type="submit" value="Send" class="form-control" /></form>
<a href="/">Go back!</a></div>
<script src="https://code.jquery.com/jquery-2.1.1.min.js"></script>
<script src="https://code.jquery.com/ui/1.11.2/jquery-ui.js"></script>
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.0/js/bootstrap.min.js"></script>
<script type="text/javascript">
var formCheckNo = 0;

$("form").submit(function( event ) {
	var form = $(this);
	var object = form.serializeArray();
	var json = JSON.stringify(object);

	if (form.attr("id") === undefined) {
		form.attr("id", "form-check-" + formCheckNo);
		formCheckNo++;
	}

	$.ajax({
		data: json,
		type: "PUT",
	}).done(function(data) {
		data = $.parseJSON(data);
		var localid = "#" + form.attr("id");

		if (data.valid) {
			$(localid + " .flash").html('<div class="alert alert-success" role="alert"></div>').children(".alert").text("Hello, " +
				object[0].value + " " + object[1].value + " " + object[2].value);
		} else {
			$(localid + " .flash").html('<div class="alert alert-danger" role="alert">Hi, There a problem with one or more of the fields!</div>');
			$(localid).effect("shake");
		}

		data.data.forEach(function(value) {
			if (value.valid) {
				if (value.warning !== "") {
					$(localid + " #form-group-" + value.count).removeClass("has-success has-error has-warning").addClass("has-warning").children(".help-block").text(value.warning);
				} else {
					$(localid + " #form-group-" + value.count).removeClass("has-success has-error has-warning").addClass("has-success").children(".help-block").text("");
				}
			} else {
				$(localid + " #form-group-" + value.count).removeClass("has-success has-error has-warning").addClass("has-error").children(".help-block").text(value.error);
			}
		});
	});

	event.preventDefault();
});
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
			flash = fmt.Sprintf(`<div class="alert alert-success" role="alert">Hello, %s %s %s!</div>`,
				es(value.Title), es(value.Firstname), es(value.Secondname))
		} else {
			flash = fmt.Sprint(`<div class="alert alert-danger" role="alert">Hi, There a problem with one or more of the fields!</div>`)
		}
		get()
	case "PUT": // For those with Javascript.
		f.ParseSerializeArray(r)
		f.MustValidate(nil, &value)
		f.Json(w)
	default:
		dir.ExecDefaultFailFn(r)
	}
}
