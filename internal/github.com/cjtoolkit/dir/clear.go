package dir

import (
	"net/http"
)

// Make sure you call this while exiting user request.
func Clear(r *http.Request) {
	mutex.Lock()
	delete(data, r)
	mutex.Unlock()
}
