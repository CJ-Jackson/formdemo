package common

import (
	html "html/template"
	"io"
	"io/ioutil"
)

func renderToHtml(reader io.Reader) (h html.HTML, i int) {
	if nil == reader {
		i = 1
		return
	}

	b, err := ioutil.ReadAll(reader)
	if nil != err {
		i = 2
		return
	}

	h = html.HTML(b)
	return
}

func ReaderToHtml(reader io.Reader) (h html.HTML) {
	h, _ = renderToHtml(reader)
	return
}
