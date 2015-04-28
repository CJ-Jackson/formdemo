package dir

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Scan Remaining Directories into variable.
func Scan(r *http.Request, a ...interface{}) (bool, error) {
	d := getDivide(r)

	if !d.setup {
		return false, fmt.Errorf("Dir: Data has not been initialise for '%p'", r)
	}

	if d.dirCount < 0 {
		return false, nil
	}

	strs := strings.Split(d.remainingPath[1:], "/")

	strs_len := len(strs)

	if strs_len < 0 {
		return false, nil
	}

	if len(a) != strs_len {
		return false, nil
	}

	for key, value := range a {
		if key >= strs_len {
			return false, nil
		}

		if value == nil {
			continue
		}

		switch t := value.(type) {
		case *string:
			*t = strs[key]
		case *int:
			n, err := strconv.ParseInt(strs[key], 10, 64)
			if err != nil {
				return false, err
			}
			*t = int(n)
		case *int64:
			n, err := strconv.ParseInt(strs[key], 10, 64)
			if err != nil {
				return false, err
			}

			*t = int64(n)
		case *int32:
			n, err := strconv.ParseInt(strs[key], 10, 32)
			if err != nil {
				return false, err
			}
			*t = int32(n)
		case *int16:
			n, err := strconv.ParseInt(strs[key], 10, 16)
			if err != nil {
				return false, err
			}
			*t = int16(n)
		case *int8:
			n, err := strconv.ParseInt(strs[key], 10, 8)
			if err != nil {
				return false, err
			}
			*t = int8(n)
		case *uint:
			n, err := strconv.ParseUint(strs[key], 10, 64)
			if err != nil {
				return false, err
			}
			*t = uint(n)
		case *uint64:
			n, err := strconv.ParseUint(strs[key], 10, 64)
			if err != nil {
				return false, err
			}
			*t = uint64(n)
		case *uint32:
			n, err := strconv.ParseUint(strs[key], 10, 32)
			if err != nil {
				return false, err
			}
			*t = uint32(n)
		case *uint16:
			n, err := strconv.ParseUint(strs[key], 10, 16)
			if err != nil {
				return false, err
			}
			*t = uint16(n)
		case *uint8:
			n, err := strconv.ParseUint(strs[key], 10, 8)
			if err != nil {
				return false, err
			}
			*t = uint8(n)
		case *float64:
			n, err := strconv.ParseFloat(strs[key], 64)
			if err != nil {
				return false, err
			}
			*t = float64(n)
		case *float32:
			n, err := strconv.ParseFloat(strs[key], 32)
			if err != nil {
				return false, err
			}
			*t = float32(n)
		}
	}
	return true, nil
}

// Scan Current Directory into variable.
func DirScan(r *http.Request, v interface{}) (err error) {
	d := getDivide(r)

	// Remove slash in front
	firstDir := d.firstDir[1:]

	switch v := v.(type) {
	case *string:
		*v = firstDir
	case *int:
		n, e := strconv.ParseInt(firstDir, 10, 64)
		err = e
		*v = int(n)
	case *int64:
		n, e := strconv.ParseInt(firstDir, 10, 64)
		err = e
		*v = int64(n)
	case *int32:
		n, e := strconv.ParseInt(firstDir, 10, 32)
		err = e
		*v = int32(n)
	case *int16:
		n, e := strconv.ParseInt(firstDir, 10, 16)
		err = e
		*v = int16(n)
	case *int8:
		n, e := strconv.ParseInt(firstDir, 10, 8)
		err = e
		*v = int8(n)
	case *uint:
		n, e := strconv.ParseUint(firstDir, 10, 64)
		err = e
		*v = uint(n)
	case *uint64:
		n, e := strconv.ParseUint(firstDir, 10, 64)
		err = e
		*v = uint64(n)
	case *uint32:
		n, e := strconv.ParseUint(firstDir, 10, 32)
		err = e
		*v = uint32(n)
	case *uint16:
		n, e := strconv.ParseUint(firstDir, 10, 16)
		err = e
		*v = uint16(n)
	case *uint8:
		n, e := strconv.ParseUint(firstDir, 10, 8)
		err = e
		*v = uint8(n)
	case *float64:
		n, e := strconv.ParseFloat(firstDir, 64)
		err = e
		*v = float64(n)
	case *float32:
		n, e := strconv.ParseFloat(firstDir, 32)
		err = e
		*v = float32(n)
	}

	return
}
