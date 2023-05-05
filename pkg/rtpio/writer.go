// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package rtpio

import (
	"io"

	"github.com/pion/rtcp"
	"github.com/pion/rtp/v2"
)

// RTPWriter is used by Interceptor.BindLocalStream.
type RTPWriter interface {
	WriteRTP(pkt *rtp.Packet) error
}

type RTPReaderFrom interface {
	ReadRTPFrom(r RTPReader) error
}

// RTCPWriter is used by Interceptor.BindRTCPWriter.
type RTCPWriter interface {
	WriteRTCP(pkts []rtcp.Packet) error
}

type RTCPReaderFrom interface {
	ReadRTCPFrom(r RTCPReader) error
}

// RawRTPWriter is a RTP packet writer that writes to an io.Writer`.`
type RawRTPWriter struct {
	dst io.Writer
}

// WriteRTP writes a RTP packet to the underlying writer.
func (w *RawRTPWriter) WriteRTP(pkt *rtp.Packet) error {
	b, err := pkt.Marshal()
	if err != nil {
		return err
	}
	_, err = w.dst.Write(b)
	return err
}

// NewRTPWriter creates a new RTP packet writer.
func NewRTPWriter(w io.Writer) RTPWriter {
	return &RawRTPWriter{
		dst: w,
	}
}

var _ RTPWriter = (*RawRTPWriter)(nil)

// RawRTCPWriter is a writer that writes RTCP packets to an `io.Writer“.
type RawRTCPWriter struct {
	dst io.Writer
}

// WriteRTCP writes a slice of RTCP packets to the underlying writer.
func (w *RawRTCPWriter) WriteRTCP(pkts []rtcp.Packet) error {
	b, err := rtcp.Marshal(pkts)
	if err != nil {
		return err
	}
	_, err = w.dst.Write(b)
	return err
}

// NewRTCPWriter creates a new RTCP packet writer.
func NewRTCPWriter(w io.Writer) RTCPWriter {
	return &RawRTCPWriter{
		dst: w,
	}
}

var _ RTCPWriter = (*RawRTCPWriter)(nil)

var DiscardRTP = discardRTPWriter{}

type discardRTPWriter struct{}

func (w discardRTPWriter) WriteRTP(pkt *rtp.Packet) error {
	return nil
}

func (w discardRTPWriter) ReadRTPFrom(r RTPReader) error {
	for {
		if _, err := r.ReadRTP(); err != nil {
			return err
		}
	}
}

var (
	_ RTPWriter     = DiscardRTP
	_ RTPReaderFrom = DiscardRTP
)

var DiscardRTCP = discardRTCPWriter{}

type discardRTCPWriter struct{}

func (w discardRTCPWriter) WriteRTCP(pkts []rtcp.Packet) error {
	return nil
}

func (w discardRTCPWriter) ReadRTCPFrom(r RTCPReader) error {
	for {
		if _, err := r.ReadRTCP(); err != nil {
			return err
		}
	}
}

var (
	_ RTCPWriter     = DiscardRTCP
	_ RTCPReaderFrom = DiscardRTCP
)

type unmarshallingRTPWriter struct {
	RTPWriter
}

// Write unmarshals the RTP packets from RTPWriter.
func (w *unmarshallingRTPWriter) Write(buf []byte) (int, error) {
	p := &rtp.Packet{}
	if err := p.Unmarshal(buf); err != nil {
		return 0, err
	}
	return len(buf), w.WriteRTP(p)
}

// NewUnmarshallingRTPWriter creates an io.Writer that Writes RTP packets from an RTPWriter.
func NewUnmarshallingRTPWriter(r RTPWriter) io.Writer {
	return &unmarshallingRTPWriter{RTPWriter: r}
}

var _ io.Writer = (*unmarshallingRTPWriter)(nil)

type unmarshallingRTCPWriter struct {
	RTCPWriter
}

// Write unmarshals the RTCP packets from RTCPWriter
func (w *unmarshallingRTCPWriter) Write(buf []byte) (int, error) {
	pkts, err := rtcp.Unmarshal(buf)
	if err != nil {
		return 0, err
	}
	return len(buf), w.WriteRTCP(pkts)
}

// NewUnmarhsallingRTCPWriter creates an io.Writer that Writes RTCP packets from an RTCPWriter.
func NewUnmarshallingRTCPWriter(r RTCPWriter) io.Writer {
	return &unmarshallingRTCPWriter{RTCPWriter: r}
}

var _ io.Writer = (*unmarshallingRTCPWriter)(nil)
