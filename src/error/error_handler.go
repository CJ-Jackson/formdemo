package error

type ErrorHandler struct {
	error ErrorInterface
	recv  interface{}
}

func NewFormDemoErrorHandler(recv interface{}) ErrorHandler {
	return ErrorHandler{
		error: GetFormDemoError(),
		recv:  recv,
	}
}
