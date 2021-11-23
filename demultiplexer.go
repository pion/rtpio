package rtpio

import (
	"io"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
)

// NewRTPRTCPDemultiplexer creates a new RFC 5761 demultiplexer.
func NewRTPRTCPDemultiplexer(r io.Reader, mtu int) (RTPReader, RTCPReader) {
	// it's ok that these are unbuffered since our API is pull-based.
	rtpReader, rtpWriter := RTPPipe()
	rtcpReader, rtcpWriter := RTCPPipe()
	go func() {
		for {
			buf := make([]byte, mtu)
			n, err := r.Read(buf)
			if err != nil {
				return
			}
			h := &rtcp.Header{}
			if err := h.Unmarshal(buf[:n]); err != nil {
				// not a valid rtp/rtcp packet.
				continue
			}
			if h.Type >= 200 && h.Type <= 207 {
				// it's an rtcp packet.
				cp, err := rtcp.Unmarshal(buf[:n])
				if err != nil {
					// not a valid rtcp packet.
					continue
				}
				if _, err := rtcpWriter.WriteRTCP(cp); err != nil {
					continue
				}
			} else {
				p := &rtp.Packet{}
				if err := p.Unmarshal(buf[:n]); err != nil {
					// not a valid rtp/rtcp packet.
					continue
				}
				if _, err := rtpWriter.WriteRTP(p); err != nil {
					continue
				}
			}
		}
	}()
	return rtpReader, rtcpReader
}
