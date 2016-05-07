package http_error

import (
	"fmt"
	"github.com/CJ-Jackson/formdemo/src/skeleton"
	"net/http"
	"strings"
)

type Error struct {
	skeleton skeleton.SkeletonInterface
}

func (e *Error) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	e.skeleton.SetResponseWriter(w)
	e.skeleton.Execute()
}

func (e *Error) setCode(code int) {
	e.skeleton.SetStatus(code)
	e.skeleton.SetTitle(fmt.Sprintf("%d: %s", code, http.StatusText(code)))
	e.skeleton.SetBody(strings.NewReader(
		fmt.Sprintf(`<h1>%d: %s</h1>`, code, http.StatusText(code))))
}

func (e *Error) NotFound() http.Handler {
	e.setCode(http.StatusNotFound)

	return e
}

func (e *Error) RaiseNotFound() {
	panic(e.NotFound())
}

func (e *Error) MethodNotAllowed() http.Handler {
	e.setCode(http.StatusMethodNotAllowed)

	return e
}

func (e *Error) RaiseMethodNotAllowed() {
	panic(e.MethodNotAllowed())
}
