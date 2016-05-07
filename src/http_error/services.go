package http_error

import (
	"github.com/CJ-Jackson/formdemo/src/skeleton"
)

func GetFormDemoError() *Error {
	return &Error{
		skeleton: skeleton.GetFormDemoSkeleton(),
	}
}
