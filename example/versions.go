package main

import (
	"log"

	"github.com/wangcong791510428/goav/avcodec"
	"github.com/wangcong791510428/goav/avdevice"
	"github.com/wangcong791510428/goav/avfilter"
	"github.com/wangcong791510428/goav/avformat"
	"github.com/wangcong791510428/goav/avutil"
	"github.com/wangcong791510428/goav/swresample"
	"github.com/wangcong791510428/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
