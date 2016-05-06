package skeleton_mock

import (
	"github.com/CJ-Jackson/shorty/src/common"
	. "github.com/smartystreets/goconvey/convey"
	html "html/template"
	"io"
	"net/http"
)

type SkeletonMock struct {
	C C

	setStatusParamStatus chan int

	setTitleParamTitle chan string

	setBodyFetchReader chan io.Reader

	executeExpected chan bool
}

func NewSkeletonMock() *SkeletonMock {
	return &SkeletonMock{
		setStatusParamStatus: make(chan int),

		setTitleParamTitle: make(chan string),

		setBodyFetchReader: make(chan io.Reader),

		executeExpected: make(chan bool),
	}
}

func (sM *SkeletonMock) SetResponseWriter(w http.ResponseWriter) {
	// Do nothing
}

func (sM *SkeletonMock) ExpectSetStatus(expectStatus int) {
	sM.setStatusParamStatus <- expectStatus
}

func (sM *SkeletonMock) SetStatus(status int) {
	sM.C.So(status, ShouldEqual, <-sM.setStatusParamStatus)
}

func (sM *SkeletonMock) ExpectSetTitle(expectTitle string) {
	sM.setTitleParamTitle <- expectTitle
}

func (sM *SkeletonMock) SetTitle(title string) {
	sM.C.So(title, ShouldEqual, <-sM.setTitleParamTitle)
}

func (sM *SkeletonMock) ExpectSetBody() html.HTML {
	return common.ReaderToHtml(<-sM.setBodyFetchReader)
}

func (sM *SkeletonMock) SetBody(body io.Reader) {
	sM.setBodyFetchReader <- body
}

func (sM *SkeletonMock) ExpectExecute() {
	sM.executeExpected <- true
}

func (sM *SkeletonMock) Execute() {
	<-sM.executeExpected
}
