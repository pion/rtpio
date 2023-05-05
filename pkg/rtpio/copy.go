// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package rtpio

// CopyRTP copies packets between a RTPReader and RTPWriter
func CopyRTP(dst RTPWriter, src RTPReader) error {
	for {
		p, err := src.ReadRTP()
		if err != nil {
			return err
		}
		if err := dst.WriteRTP(p); err != nil {
			return err
		}
	}
}

// CopyRTCP copies packets between a RTCPReader and RTCPWriter
func CopyRTCP(dst RTCPWriter, src RTCPReader) error {
	for {
		pkts, err := src.ReadRTCP()
		if err != nil {
			return err
		}
		if err := dst.WriteRTCP(pkts); err != nil {
			return err
		}
	}
}
