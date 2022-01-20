package rtpio

import (
	"io"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
)

type pipeRTPReader struct {
	closer    io.Closer
	rtpReader RTPReader
}

func (r *pipeRTPReader) ReadRTP() (*rtp.Packet, error) {
	return r.rtpReader.ReadRTP()
}

func (r *pipeRTPReader) Close() error {
	return r.closer.Close()
}

var _ RTPReader = (*pipeRTPReader)(nil)
var _ io.Closer = (*pipeRTPReader)(nil)

type pipeRTPWriter struct {
	closer    io.Closer
	rtpWriter RTPWriter
}

func (w *pipeRTPWriter) WriteRTP(pkt *rtp.Packet) error {
	return w.rtpWriter.WriteRTP(pkt)
}

func (w *pipeRTPWriter) Close() error {
	return w.closer.Close()
}

var _ RTPWriter = (*pipeRTPWriter)(nil)
var _ io.Closer = (*pipeRTPWriter)(nil)

type pipeRTCPReader struct {
	closer     io.Closer
	rtcpReader RTCPReader
}

func (r *pipeRTCPReader) ReadRTCP() ([]rtcp.Packet, error) {
	return r.rtcpReader.ReadRTCP()
}

func (r *pipeRTCPReader) Close() error {
	return r.closer.Close()
}

var _ RTCPReader = (*pipeRTCPReader)(nil)
var _ io.Closer = (*pipeRTCPReader)(nil)

type pipeRTCPWriter struct {
	closer     io.Closer
	rtcpWriter RTCPWriter
}

func (w *pipeRTCPWriter) WriteRTCP(pkts []rtcp.Packet) error {
	return w.rtcpWriter.WriteRTCP(pkts)
}

func (w *pipeRTCPWriter) Close() error {
	return w.closer.Close()
}

var _ RTCPWriter = (*pipeRTCPWriter)(nil)
var _ io.Closer = (*pipeRTCPWriter)(nil)

// RTPPipe creates a new RTPPipe and returns the reader and writer.
func RTPPipe() (RTPReadCloser, RTPWriteCloser) {
	r, w := io.Pipe()
	return &pipeRTPReader{closer: r, rtpReader: NewRTPReader(r, 1500)}, &pipeRTPWriter{closer: w, rtpWriter: NewRTPWriter(w)}
}

// RTCPPipe creates a new RTCPPipe and returns the reader and writer.
func RTCPPipe() (RTCPReadCloser, RTCPWriteCloser) {
	r, w := io.Pipe()
	return &pipeRTCPReader{closer: r, rtcpReader: NewRTCPReader(r, 1500)}, &pipeRTCPWriter{closer: w, rtcpWriter: NewRTCPWriter(w)}
}
