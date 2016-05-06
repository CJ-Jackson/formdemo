package skeleton

import "net/http"

func GetFormDemoSkeleton() *Skeleton {
	return &Skeleton{
		tpl:    skeletonHtml,
		status: http.StatusOK,
	}
}
