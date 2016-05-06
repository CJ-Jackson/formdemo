package skeleton

import (
	"io"
	"net/http"
)

type SkeletonInterface interface {
	SetResponseWriter(w http.ResponseWriter)
	SetStatus(status int)
	SetTitle(title string)
	SetHead(head io.Reader)
	SetBody(body io.Reader)
	SetFooter(footer io.Reader)
	SetJavascript(javascript io.Reader)
	Execute()
}
