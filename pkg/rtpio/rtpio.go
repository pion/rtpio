// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package rtpio implements `io.Reader` and `io.Writer` style interfaces for RTP and RTCP packets.
package rtpio

import "io"

type RTPReadCloser interface {
	RTPReader
	io.Closer
}

type RTPWriteCloser interface {
	RTPWriter
	io.Closer
}

type RTPReadWriter interface {
	RTPReader
	RTPWriter
}

type RTPReadWriteCloser interface {
	RTPReader
	RTPWriter
	io.Closer
}

type RTCPReadCloser interface {
	RTCPReader
	io.Closer
}

type RTCPWriteCloser interface {
	RTCPWriter
	io.Closer
}

type RTCPReadWriter interface {
	RTCPReader
	RTCPWriter
}

type RTCPReadWriteCloser interface {
	RTCPReader
	RTCPWriter
	io.Closer
}
