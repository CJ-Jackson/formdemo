package http_error

import "net/http"

type ErrorHandler struct {
	recv interface{}
}

func NewFormDemoErrorHandler(recv interface{}) ErrorHandler {
	return ErrorHandler{
		recv: recv,
	}
}

func (e ErrorHandler) ServeHTTP(_ http.ResponseWriter, _ *http.Request) {
	// Do nothing
}

func (e ErrorHandler) getHttpHandler() (http.Handler, int) {
	switch value := e.recv.(type) {
	case nil:
		// Do nothing
	case http.Handler:
		return value, 1
	}

	return e, 0
}

func (e ErrorHandler) GetHttpHandler() http.Handler {
	h, _ := e.getHttpHandler()
	return h
}
