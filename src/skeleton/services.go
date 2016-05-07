package skeleton

import "net/http"

func GetFormDemoSkeleton(options ...func(*Skeleton)) *Skeleton {
	s := &Skeleton{
		tpl:    skeletonHtml,
		status: http.StatusOK,
	}

	for _, option := range options {
		option(s)
	}

	return s
}
