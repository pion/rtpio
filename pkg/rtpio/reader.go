package rtpio

import (
	"io"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
)

// RTPReader is used by Interceptor.BindRemoteStream.
type RTPReader interface {
	ReadRTP() (*rtp.Packet, error)
}

type RTPWriterTo interface {
	WriteRTPTo(w RTPWriter) error
}

// RTCPReader is used by Interceptor.BindRTCPReader.
type RTCPReader interface {
	// Read a batch of rtcp packets. This returns the number of packets read, not the number of bytes!
	ReadRTCP() ([]rtcp.Packet, error)
}

type RTCPWriterTo interface {
	WriteRTCPTo(w RTCPWriter) error
}

// RawRTPReader is a RTPReader that reads from an `io.Reader`.
type RawRTPReader struct {
	src io.Reader
	mtu int
}

// ReadRTP reads a single RTP packet from the underlying reader.
func (r *RawRTPReader) ReadRTP() (*rtp.Packet, error) {
	buf := make([]byte, r.mtu)
	n, err := r.src.Read(buf)
	if err != nil {
		return nil, err
	}
	pkt := &rtp.Packet{}
	if err := pkt.Unmarshal(buf[:n]); err != nil {
		return nil, err
	}
	return pkt, nil
}

// NewRTPReader creates a new RTP packet reader.
func NewRTPReader(r io.Reader, mtu int) RTPReader {
	return &RawRTPReader{src: r, mtu: mtu}
}

var _ RTPReader = (*RawRTPReader)(nil)

// RawRTCPReader is a RTCPReader that reads from an `io.Reader`.
type RawRTCPReader struct {
	src io.Reader
	mtu int
}

// ReadRTCP reads a batch of RTCP packets from the underlying reader.
func (r *RawRTCPReader) ReadRTCP() ([]rtcp.Packet, error) {
	// read from backlog first.
	buf := make([]byte, r.mtu)
	n, err := r.src.Read(buf)
	if err != nil {
		return nil, err
	}
	return rtcp.Unmarshal(buf[:n])
}

// NewRTCPReader creates a new RTCP packet reader.
func NewRTCPReader(r io.Reader, mtu int) RTCPReader {
	return &RawRTCPReader{src: r, mtu: mtu}
}

var _ RTCPReader = (*RawRTCPReader)(nil)
