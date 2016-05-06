package error

type ErrorHandler struct {
	error ErrorInterface
	recv  interface{}
}
