// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package rtpiotest provides helpers for testing the rtpio package
package rtpiotest

import (
	"log"

	"github.com/pion/rtcp"
	"github.com/pion/rtp/v2"
	"github.com/pion/rtpio/pkg/rtpio"
)

type writeRTPLogger struct {
	prefix string
	w      rtpio.RTPWriter
}

func (l *writeRTPLogger) WriteRTP(p *rtp.Packet) (err error) {
	err = l.w.WriteRTP(p)
	if err != nil {
		log.Printf("%s (err: %v) %v", l.prefix, err, p) //nolint:forbidigo
	} else {
		log.Printf("%s %v", l.prefix, p) //nolint:forbidigo
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
		log.Printf("%s (err: %v) %v", l.prefix, err, p) //nolint:forbidigo
	} else {
		log.Printf("%s %v", l.prefix, p) //nolint:forbidigo
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
		log.Printf("%s (err: %v) %v", l.prefix, err, p) //nolint:forbidigo
	} else {
		log.Printf("%s %v", l.prefix, p) //nolint:forbidigo
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
		log.Printf("%s (err: %v) %v", l.prefix, err, p) //nolint:forbidigo
	} else {
		log.Printf("%s %x", l.prefix, p) //nolint:forbidigo
	}
	return
}

// NewReadRTCPLogger returns a reader that behaves like r except
// that it logs (using log.Printf) each read to standard error,
// printing the prefix and the hexadecimal data read.
func NewReadRTCPLogger(prefix string, r rtpio.RTCPReader) rtpio.RTCPReader {
	return &readRTCPLogger{prefix, r}
}
