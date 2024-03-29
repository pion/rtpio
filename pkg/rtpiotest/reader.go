// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package rtpiotest

import (
	"testing/iotest"

	"github.com/pion/rtpio/pkg/rtpio"
)

// DataErrRTPReader is ..
func DataErrRTPReader(r rtpio.RTPReader) rtpio.RTPReader {
	return rtpio.NewRTPReader(iotest.DataErrReader(rtpio.NewUnmarshallingRTPReader(r)), 1500)
}

// ErrRTPReader is ..
func ErrRTPReader(err error) rtpio.RTPReader {
	return rtpio.NewRTPReader(iotest.ErrReader(err), 1500)
}

// HalfRTPReader is ..
func HalfRTPReader(r rtpio.RTPReader) rtpio.RTPReader {
	return rtpio.NewRTPReader(iotest.HalfReader(rtpio.NewUnmarshallingRTPReader(r)), 1500)
}

// TimeoutRTPReader is ..
func TimeoutRTPReader(r rtpio.RTPReader) rtpio.RTPReader {
	return rtpio.NewRTPReader(iotest.TimeoutReader(rtpio.NewUnmarshallingRTPReader(r)), 1500)
}

// DataErrRTCPReader is ..
func DataErrRTCPReader(r rtpio.RTCPReader) rtpio.RTCPReader {
	return rtpio.NewRTCPReader(iotest.DataErrReader(rtpio.NewUnmarshallingRTCPReader(r)), 1500)
}

// ErrRTCPReader is ..
func ErrRTCPReader(err error) rtpio.RTCPReader {
	return rtpio.NewRTCPReader(iotest.ErrReader(err), 1500)
}

// HalfRTCPReader is ..
func HalfRTCPReader(r rtpio.RTCPReader) rtpio.RTCPReader {
	return rtpio.NewRTCPReader(iotest.HalfReader(rtpio.NewUnmarshallingRTCPReader(r)), 1500)
}

// TimeoutRTCPReader is ..
func TimeoutRTCPReader(r rtpio.RTCPReader) rtpio.RTCPReader {
	return rtpio.NewRTCPReader(iotest.TimeoutReader(rtpio.NewUnmarshallingRTCPReader(r)), 1500)
}
