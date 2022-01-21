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

// rawRTPReader is a RTPReader that reads from an `io.Reader`.
type rawRTPReader struct {
	src io.Reader
	mtu int
}

// ReadRTP reads a single RTP packet from the underlying reader.
func (r *rawRTPReader) ReadRTP() (*rtp.Packet, error) {
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
func NewRTPReader(r io.Reader, mtu int) *rawRTPReader {
	return &rawRTPReader{src: r, mtu: mtu}
}

var _ RTPReader = (*rawRTPReader)(nil)

// rawRTCPReader is a RTCPReader that reads from an `io.Reader`.
type rawRTCPReader struct {
	src io.Reader
	mtu int
}

// ReadRTCP reads a batch of RTCP packets from the underlying reader.
func (r *rawRTCPReader) ReadRTCP() ([]rtcp.Packet, error) {
	// read from backlog first.
	buf := make([]byte, r.mtu)
	n, err := r.src.Read(buf)
	if err != nil {
		return nil, err
	}
	return rtcp.Unmarshal(buf[:n])
}

// NewRTCPReader creates a new RTCP packet reader.
func NewRTCPReader(r io.Reader, mtu int) *rawRTCPReader {
	return &rawRTCPReader{src: r, mtu: mtu}
}

var _ RTCPReader = (*rawRTCPReader)(nil)

type unmarshallingRTPReader struct {
	RTPReader
}

// Read unmarshals the RTP packets from RTPReader.
func (r *unmarshallingRTPReader) Read(p []byte) (int, error) {
	pkt, err := r.ReadRTP()
	if err != nil {
		return 0, err
	}
	return pkt.MarshalTo(p)
}

// NewUnmarshallingRTPReader creates an io.Reader that reads RTP packets from an RTPReader.
func NewUnmarshallingRTPReader(r RTPReader) io.Reader {
	return &unmarshallingRTPReader{RTPReader: r}
}

var _ io.Reader = (*unmarshallingRTPReader)(nil)

type unmarshallingRTCPReader struct {
	RTCPReader
}

// Read unmarshals the RTCP packets from RTCPReader
func (r *unmarshallingRTCPReader) Read(p []byte) (int, error) {
	pkts, err := r.ReadRTCP()
	if err != nil {
		return 0, err
	}
	buf, err := rtcp.Marshal(pkts)
	if err != nil {
		return 0, err
	}
	if len(buf) < len(p) {
		return 0, io.ErrShortBuffer
	}
	return copy(p, buf), nil
}

// NewUnmarhsallingRTCPReader creates an io.Reader that reads RTCP packets from an RTCPReader.
func NewUnmarshallingRTCPReader(r RTCPReader) io.Reader {
	return &unmarshallingRTCPReader{RTCPReader: r}
}

var _ io.Reader = (*unmarshallingRTCPReader)(nil)