// URL/Dir Path Toolkit
package dir

import (
	"net/http"
	"strings"
)

// Clean up path.
func cleanPath(p string) string {
	if p == "" {
		return "/"
	}

	n := len(p)
	var buf []byte

	r := 1
	w := 1

	if p[0] != '/' {
		r = 0
		buf = make([]byte, n+1)
		buf[0] = '/'
	}

	trailing := n > 2 && p[n-1] == '/'

	for r < n {
		switch {
		case p[r] == '/':
			r++

		case p[r] == '.' && r+1 == n:
			trailing = true
			r++

		case p[r] == '.' && p[r+1] == '/':
			r++

		case p[r] == '.' && p[r+1] == '.' && (r+2 == n || p[r+2] == '/'):
			r += 2

			if w > 1 {
				w--

				if buf == nil {
					for w > 1 && p[w] != '/' {
						w--
					}
				} else {
					for w > 1 && buf[w] != '/' {
						w--
					}
				}
			}

		default:
			if w > 1 {
				lazyBuf(&buf, p, w, '/')
				w++
			}

			for r < n && p[r] != '/' {
				lazyBuf(&buf, p, w, p[r])
				w++
				r++
			}
		}
	}

	if trailing && w > 1 {
		lazyBuf(&buf, p, w, '/')
		w++
	}

	if buf == nil {
		return p[:w]
	}
	return string(buf[:w])
}

func lazyBuf(buf *[]byte, s string, w int, c byte) {
	if *buf == nil {
		if s[w] == c {
			return
		}

		*buf = make([]byte, len(s))
		copy(*buf, s[:w])
	}
	(*buf)[w] = c
}

func _divide(path string) (firstDir, remainingPath string) {
	if path == "/" || path == "" {
		firstDir = "/"
		return
	}

	if path[0] == '/' {
		firstDir += "/"
		path = path[1:]
	}

	index := strings.Index(path, "/")
	if index == -1 {
		firstDir += path
		return
	}

	firstDir += path[:index]
	remainingPath = path[index:]
	return
}

// Divide Path by first directory, return path of first directory and remaining directory
func Divide(path string) (firstDir, remainingPath string) {
	path = cleanPath(path)
	firstDir, remainingPath = _divide(path)
	return
}

// As Divide but with path converted to lower case.
func DivideLower(path string) (firstDir, remainingPath string) {
	firstDir, remainingPath = Divide(strings.ToLower(path))
	return
}

// Divide Path by first directory and return the first directory from user http request.
// CurrentPath and RemainingDirectory will be stored temporarly in memory, can be access
// by calling FetchData(r *http.Request)
func DivideHttpPath(r *http.Request) (firstDir string) {
	if r == nil {
		return
	}
	var path string

	d := getDivide(r)

	if !d.setup {
		d.setup = true
		path = r.URL.Path
		path = cleanPath(path)
		// First directory does not count, so subtract 1.
		d.dirCount = strings.Count(path, "/") - 1
	} else {
		path = d.remainingPath
	}

	firstDir, d.remainingPath = _divide(path)

	d.firstDir = firstDir
	d.currentPath += firstDir
	d.dirCount--

	return
}

// Execute passFn if Remaining Path is Blank, else execute failFn.
// To use default failFn, specify 'failFn' as 'nil'
func B(r *http.Request, passFn, failFn func()) {
	if r == nil {
		return
	}

	d := getDivide(r)

	if !d.setup {
		return
	}

	if d.dirCount < 0 {
		passFn()
	} else if failFn != nil {
		failFn()
	} else if d.failFn != nil {
		d.failFn()
	}
}

// Execute failFn if dirDepth does not match what remaining else execute passFn.
// To use default failFn, specify 'failFn' as 'nil'
func L(dirDepth int, r *http.Request, passFn, failFn func()) {
	if r == nil {
		return
	}

	d := getDivide(r)

	if !d.setup {
		return
	}

	if (d.dirCount + 1) == dirDepth {
		passFn()
	} else {
		if failFn != nil {
			failFn()
		} else if d.failFn != nil {
			d.failFn()
		}
	}
}

// Advanced version of L.
func LMap(r *http.Request, passFnMap map[int]func(), failFn func()) {
	if r == nil {
		return
	}

	d := getDivide(r)

	if !d.setup {
		return
	}

	switch {
	case passFnMap == nil:
		fallthrough
	default:
		if failFn != nil {
			failFn()
		} else if d.failFn != nil {
			d.failFn()
		}
	case passFnMap[d.dirCount+1] != nil:
		passFnMap[d.dirCount+1]()
	}
}

// Set Default failFn
func SetDefaultFailFn(r *http.Request, failFn func()) {
	if r == nil {
		return
	}
	d := getDivide(r)
	d.failFn = failFn
}

// Exec Default failFn
func ExecDefaultFailFn(r *http.Request) {
	if r == nil {
		return
	}
	d := getDivide(r)
	if d.failFn != nil {
		d.failFn()
	}
}
