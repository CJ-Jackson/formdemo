package skeleton

import (
	"io"
	"net/http"
)

type SkeletonInterface interface {
	SetResponseWriter(w http.ResponseWriter)
	SetStatus(status int)
	SetTitle(title string)
	SetBody(body io.Reader)
	Execute()
}
