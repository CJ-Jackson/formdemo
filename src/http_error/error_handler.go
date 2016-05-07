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

func (e ErrorHandler) getHttpHandler() (http.Handler, int) {
	switch value := e.recv.(type) {
	case nil:
		// Do nothing
	case http.Handler:
		return value, 1
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}), 0
}

func (e ErrorHandler) GetHttpHandler() http.Handler {
	h, _ := e.getHttpHandler()
	return h
}
