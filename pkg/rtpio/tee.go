package rtpio

import (
	"github.com/pion/rtcp"
	"github.com/pion/rtp"
)

type teeRTPReader struct {
	r RTPReader
	w RTPWriter
}

func (r *teeRTPReader) ReadRTP() (*rtp.Packet, error) {
	pkt, err := r.r.ReadRTP()
	if err != nil {
		return nil, err
	}
	if err := r.w.WriteRTP(pkt); err != nil {
		return nil, err
	}
	return pkt, nil
}

func TeeRTPReader(r RTPReader, w RTPWriter) RTPReader {
	return &teeRTPReader{r: r, w: w}
}

type teeRTCPReader struct {
	r RTCPReader
	w RTCPWriter
}

func (r *teeRTCPReader) ReadRTCP() ([]rtcp.Packet, error) {
	pkts, err := r.r.ReadRTCP()
	if err != nil {
		return nil, err
	}
	if err := r.w.WriteRTCP(pkts); err != nil {
		return nil, err
	}
	return pkts, nil
}

func TeeRTCPReader(r RTCPReader, w RTCPWriter) RTCPReader {
	return &teeRTCPReader{r: r, w: w}
}