// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avutil
/*
	#cgo pkg-config: libavutil
	#include <libavutil/dict.h>
	#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

func AvDictSet(dt **Dictionary, key, value string, flags int) int {
	return int(C.av_dict_set((**C.struct_AVDictionary)(unsafe.Pointer(dt)), C.CString(key), C.CString(value), C.int(flags)))
}
