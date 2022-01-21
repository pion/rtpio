package rtpiotest

import (
	"log"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/rtpio/pkg/rtpio"
)

type writeRTPLogger struct {
	prefix string
	w      rtpio.RTPWriter
}

func (l *writeRTPLogger) WriteRTP(p *rtp.Packet) (err error) {
	err = l.w.WriteRTP(p)
	if err != nil {
		log.Printf("%s %x: %v", l.prefix, p, err)
	} else {
		log.Printf("%s %x", l.prefix, p)
	}
	return
}

// NewWriteRTPLogger returns a writer that behaves like w except
// that it logs (using log.Printf) each write to standard error,
// printing the prefix and the hexadecimal data written.
func NewWriteRTPLogger(prefix string, w rtpio.RTPWriter) rtpio.RTPWriter {
	return &writeRTPLogger{prefix, w}
}

type readRTPLogger struct {
	prefix string
	r      rtpio.RTPReader
}

func (l *readRTPLogger) ReadRTP() (p *rtp.Packet, err error) {
	p, err = l.r.ReadRTP()
	if err != nil {
		log.Printf("%s %x: %v", l.prefix, p, err)
	} else {
		log.Printf("%s %x", l.prefix, p)
	}
	return
}

// NewReadRTPLogger returns a reader that behaves like r except
// that it logs (using log.Printf) each read to standard error,
// printing the prefix and the hexadecimal data read.
func NewReadRTPLogger(prefix string, r rtpio.RTPReader) rtpio.RTPReader {
	return &readRTPLogger{prefix, r}
}

type writeRTCPLogger struct {
	prefix string
	w      rtpio.RTCPWriter
}

func (l *writeRTCPLogger) WriteRTCP(p []rtcp.Packet) (err error) {
	err = l.w.WriteRTCP(p)
	if err != nil {
		log.Printf("%s %x: %v", l.prefix, p, err)
	} else {
		log.Printf("%s %x", l.prefix, p)
	}
	return
}

// NewWriteRTCPLogger returns a writer that behaves like w except
// that it logs (using log.Printf) each write to standard error,
// printing the prefix and the hexadecimal data written.
func NewWriteRTCPLogger(prefix string, w rtpio.RTCPWriter) rtpio.RTCPWriter {
	return &writeRTCPLogger{prefix, w}
}

type readRTCPLogger struct {
	prefix string
	r      rtpio.RTCPReader
}

func (l *readRTCPLogger) ReadRTCP() (p []rtcp.Packet, err error) {
	p, err = l.r.ReadRTCP()
	if err != nil {
		log.Printf("%s %x: %v", l.prefix, p, err)
	} else {
		log.Printf("%s %x", l.prefix, p)
	}
	return
}

// NewReadRTCPLogger returns a reader that behaves like r except
// that it logs (using log.Printf) each read to standard error,
// printing the prefix and the hexadecimal data read.
func NewReadRTCPLogger(prefix string, r rtpio.RTCPReader) rtpio.RTCPReader {
	return &readRTCPLogger{prefix, r}
}