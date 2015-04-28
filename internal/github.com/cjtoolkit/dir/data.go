package dir

import (
	"net/http"
	"sync"
)

type divide struct {
	setup         bool
	firstDir      string
	currentPath   string
	remainingPath string
	dirCount      int
	failFn        func()
}

func getDivide(r *http.Request) *divide {
	mutex.RLock()
	d := data[r]
	mutex.RUnlock()
	if d == nil {
		d = &divide{}
		mutex.Lock()
		data[r] = d
		mutex.Unlock()
	}
	return d
}

var (
	mutex sync.RWMutex
	data  = map[*http.Request]*divide{}
)

// Get Path Data, return Current Path and Remaining Path
func FetchData(r *http.Request) (currentPath, remainingPath string) {
	d := getDivide(r)
	currentPath = d.currentPath
	remainingPath = d.remainingPath
	return
}
