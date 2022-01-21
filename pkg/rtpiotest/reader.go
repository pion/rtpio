package rtpiotest

import (
	"testing/iotest"

	"github.com/muxable/rtpio/pkg/rtpio"
)

func DataErrRTPReader(r rtpio.RTPReader) rtpio.RTPReader {
	return rtpio.NewRTPReader(iotest.DataErrReader(rtpio.NewUnmarshallingRTPReader(r)), 1500)
}

func ErrRTPReader(err error) rtpio.RTPReader {
	return rtpio.NewRTPReader(iotest.ErrReader(err), 1500)
}

func HalfRTPReader(r rtpio.RTPReader) rtpio.RTPReader {
	return rtpio.NewRTPReader(iotest.HalfReader(rtpio.NewUnmarshallingRTPReader(r)), 1500)
}

func NewRTPReadLogger(prefix string, r rtpio.RTPReader) rtpio.RTPReader {
	return rtpio.NewRTPReader(iotest.NewReadLogger(prefix, rtpio.NewUnmarshallingRTPReader(r)), 1500)
}

func NewRTPWriteLogger(prefix string, r rtpio.RTPWriter) rtpio.RTPWriter {
	return rtpio.NewRTPWriter(iotest.NewWriteLogger(prefix, rtpio.NewUnmarshallingRTPWriter(r)))
}

func TimeoutRTPReader(r rtpio.RTPReader) rtpio.RTPReader {
	return rtpio.NewRTPReader(iotest.TimeoutReader(rtpio.NewUnmarshallingRTPReader(r)), 1500)
}

func DataErrRTCPReader(r rtpio.RTCPReader) rtpio.RTCPReader {
	return rtpio.NewRTCPReader(iotest.DataErrReader(rtpio.NewUnmarshallingRTCPReader(r)), 1500)
}

func ErrRTCPReader(err error) rtpio.RTCPReader {
	return rtpio.NewRTCPReader(iotest.ErrReader(err), 1500)
}

func HalfRTCPReader(r rtpio.RTCPReader) rtpio.RTCPReader {
	return rtpio.NewRTCPReader(iotest.HalfReader(rtpio.NewUnmarshallingRTCPReader(r)), 1500)
}

func NewRTCPReadLogger(prefix string, r rtpio.RTCPReader) rtpio.RTCPReader {
	return rtpio.NewRTCPReader(iotest.NewReadLogger(prefix, rtpio.NewUnmarshallingRTCPReader(r)), 1500)
}

func NewRTCPWriteLogger(prefix string, r rtpio.RTCPWriter) rtpio.RTCPWriter {
	return rtpio.NewRTCPWriter(iotest.NewWriteLogger(prefix, rtpio.NewUnmarshallingRTCPWriter(r)))
}

func TimeoutRTCPReader(r rtpio.RTCPReader) rtpio.RTCPReader {
	return rtpio.NewRTCPReader(iotest.TimeoutReader(rtpio.NewUnmarshallingRTCPReader(r)), 1500)
}
