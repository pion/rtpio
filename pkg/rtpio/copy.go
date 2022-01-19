package rtpio

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
