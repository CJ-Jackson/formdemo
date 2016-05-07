package skeleton

import (
	"github.com/CJ-Jackson/formdemo/src/common"
	html "html/template"
	"io"
	"net/http"
)

type Skeleton struct {
	tpl    *html.Template
	w      http.ResponseWriter
	status int
	title  string
	body   io.Reader
}

func (s *Skeleton) SetResponseWriter(w http.ResponseWriter) {
	s.w = w
}

func (s *Skeleton) SetStatus(status int) {
	s.status = status
}

func (s *Skeleton) SetTitle(title string) {
	s.title = title
}

func (s *Skeleton) Title() string {
	return s.title
}

func (s *Skeleton) SetBody(body io.Reader) {
	s.body = body
}

func (s *Skeleton) Body() html.HTML {
	return common.ReaderToHtml(s.body)
}

func (s *Skeleton) execute(wr io.Writer) {
	s.tpl.Execute(wr, s)
}

func (s *Skeleton) Execute() {
	s.w.Header().Set("Content-Type", "text/html; charset=utf-8")
	s.w.WriteHeader(s.status)
	s.execute(s.w)
}

func SkeletonResponseWriter(w http.ResponseWriter) func(*Skeleton) {
	return func(s *Skeleton) {
		s.w = w
	}
}
