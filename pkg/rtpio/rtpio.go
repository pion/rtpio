// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package rtpio implements `io.Reader` and `io.Writer` style interfaces for RTP and RTCP packets.
package rtpio

import "io"

// RTPReadCloser is ...
type RTPReadCloser interface {
	RTPReader
	io.Closer
}

// RTPWriteCloser is ...
type RTPWriteCloser interface {
	RTPWriter
	io.Closer
}

// RTPReadWriter is ...
type RTPReadWriter interface {
	RTPReader
	RTPWriter
}

// RTPReadWriteCloser is ...
type RTPReadWriteCloser interface {
	RTPReader
	RTPWriter
	io.Closer
}

// RTCPReadCloser is ...
type RTCPReadCloser interface {
	RTCPReader
	io.Closer
}

// RTCPWriteCloser is ...
type RTCPWriteCloser interface {
	RTCPWriter
	io.Closer
}

// RTCPReadWriter is ...
type RTCPReadWriter interface {
	RTCPReader
	RTCPWriter
}

// RTCPReadWriteCloser is ...
type RTCPReadWriteCloser interface {
	RTCPReader
	RTCPWriter
	io.Closer
}
