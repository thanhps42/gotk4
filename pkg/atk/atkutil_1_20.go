// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"
)

// #include <stdlib.h>
// #include <atk/atk.h>
import "C"

// GetVersion gets the current version for ATK.
//
// The function returns the following values:
//
//    - utf8: version string for ATK.
//
func GetVersion() string {
	var _cret *C.gchar // in

	_cret = C.atk_get_version()

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
