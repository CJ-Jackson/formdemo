package error

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

func (e *Error) NotFound() http.Handler {
	e.skeleton.SetStatus(http.StatusNotFound)
	e.skeleton.SetTitle(fmt.Sprintf("%d: %s", http.StatusNotFound, http.StatusText(http.StatusNotFound)))
	e.skeleton.SetBody(strings.NewReader(
		fmt.Sprintf(`<h1>%d: %s</h1>`, http.StatusNotFound, http.StatusText(http.StatusNotFound))))

	return e
}

func (e *Error) RaiseNotFound() {
	panic(e.NotFound())
}
