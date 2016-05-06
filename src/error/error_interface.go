package error

type ErrorInterface interface {
	NotFound()
	RaiseNotFound()
}
