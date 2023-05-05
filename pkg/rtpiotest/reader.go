// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package rtpiotest

import (
	"testing/iotest"

	"github.com/pion/rtpio/pkg/rtpio"
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

func TimeoutRTCPReader(r rtpio.RTCPReader) rtpio.RTCPReader {
	return rtpio.NewRTCPReader(iotest.TimeoutReader(rtpio.NewUnmarshallingRTCPReader(r)), 1500)
}
