package main

import (
	"log"

	"github.com/trafficstars/goav/avcodec"
	"github.com/trafficstars/goav/avdevice"
	"github.com/trafficstars/goav/avfilter"
	"github.com/trafficstars/goav/avformat"
	"github.com/trafficstars/goav/avutil"
	"github.com/trafficstars/goav/swresample"
	"github.com/trafficstars/goav/swscale"
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
