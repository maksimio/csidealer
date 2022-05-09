package decoder

const (
	BITS_PER_BYTE   = 8
	BITS_PER_SYMBOL = 10
)

func bitConvert(data int) int {
	if data&512 != 0 {
		data -= 1024
	}
	return data
}

func decodeCsi(dataCsi []byte, nr, nc, numTones uint8) [][]complex128 {
	csi := make([][]complex128, nr*nc)
	for i := range csi {
		csi[i] = make([]complex128, numTones)
	}

	bitsLeft := 16
	hData := uint32(dataCsi[0]) + (uint32(dataCsi[1]) << BITS_PER_BYTE)
	current_data := hData & 65535
	idx := 2

	var k, ncIdx, nrIdx uint8
	for ; k < numTones; k++ {
		ncIdx = 0
		for ; ncIdx < nc; ncIdx++ {
			nrIdx = 0
			for ; nrIdx < nr; nrIdx++ {
				if bitsLeft < BITS_PER_SYMBOL {
					hData = uint32(dataCsi[idx]) + (uint32(dataCsi[idx+1]) << BITS_PER_BYTE)
					idx += 2
					current_data += hData << bitsLeft
					bitsLeft += 16
				}
				imag := current_data & 1023
				bitsLeft -= BITS_PER_SYMBOL
				current_data = current_data >> BITS_PER_SYMBOL

				if bitsLeft < BITS_PER_SYMBOL {
					hData = uint32(dataCsi[idx]) + (uint32(dataCsi[idx+1]) << BITS_PER_BYTE)
					idx += 2
					current_data += hData << bitsLeft
					bitsLeft += 16
				}
				real := current_data & 1023
				bitsLeft -= BITS_PER_SYMBOL
				current_data = current_data >> BITS_PER_SYMBOL

				csi[nrIdx+ncIdx*nr][k] = complex(float64(bitConvert(int(real))), float64(bitConvert(int(imag))))
			}
		}
	}
	return csi
}
