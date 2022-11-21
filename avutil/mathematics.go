package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
//#include <libavutil/mathematics.h>
//#include <stdlib.h>
import "C"
import "unsafe"
import "github.com/wangcong791510428/goav/avcodec"

//Return a string describing the media_type enum, NULL if media_type is unknown.
func AvRescaleQ(a int64, bq Rational, cq Rational) int64 {
	return int64(C.av_rescale_q(C.int64_t(a), (C.struct_AVRational)(bq),  (C.struct_AVRational)(cq)))
}

func toCRational(rational *avcodec.Rational) *C.struct_AVRational {
	return (*C.struct_AVRational)(unsafe.Pointer(rational))
}
