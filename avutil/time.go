package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
//#include <libavutil/time.h>
//#include <stdlib.h>
//#include <stdint.h>
import "C"

func AvGettime() int64 {
	return int64(C.av_gettime())
}

func AvUsleep(aaa int) int {
	return int(C.av_usleep(C.uint(aaa)))
}