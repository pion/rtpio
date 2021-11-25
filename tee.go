package rtpio

import (
	"github.com/pion/rtcp"
	"github.com/pion/rtp"
)

type teeRTPReader struct {
	r RTPReader
	w RTPWriter
}

func (r *teeRTPReader) ReadRTP(pkt *rtp.Packet) (int, error) {
	n, err := r.r.ReadRTP(pkt)
	if err != nil {
		return n, err
	}
	return r.w.WriteRTP(pkt)
}

func TeeRTPReader(r RTPReader, w RTPWriter) RTPReader {
	return &teeRTPReader{r: r, w: w}
}

type teeRTCPReader struct {
	r RTCPReader
	w RTCPWriter
}

func (r *teeRTCPReader) ReadRTCP(pkts []rtcp.Packet) (int, error) {
	n, err := r.r.ReadRTCP(pkts)
	if err != nil {
		return n, err
	}
	return r.w.WriteRTCP(pkts)
}

func TeeRTCPReader(r RTCPReader, w RTCPWriter) RTCPReader {
	return &teeRTCPReader{r: r, w: w}
}
