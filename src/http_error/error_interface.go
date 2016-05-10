package http_error

type ErrorInterface interface {
	NotFound()
	RaiseNotFound()
	MethodNotAllowed()
	RaiseMethodNotAllowed()
}
