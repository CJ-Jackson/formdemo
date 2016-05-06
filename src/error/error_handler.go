package error

type ErrorHandler struct {
	error ErrorInterface
	recv  interface{}
}

func NewErrorHandler(recv interface{}) ErrorHandler {
	return ErrorHandler{
		error: GetFormDemoError(),
		recv:  recv,
	}
}
