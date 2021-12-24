package csicore

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

func DecodeCsi(local_h []byte, nr, nc, numTones uint8) [][]complex128 {
	csi := make([][]complex128, nr*nc)
	for i := range csi {
		csi[i] = make([]complex128, numTones)
	}

	bits_left := 16
	h_data := uint32(local_h[0]) + (uint32(local_h[1]) << BITS_PER_BYTE)
	current_data := h_data & 65535
	idx := 2

	var k, nc_idx, nr_idx uint8 = 0, 0, 0
	for ; k < numTones; k++ {
		for ; nc_idx < nc; nc_idx++ {
			for ; nr_idx < nc; nr_idx++ {
				if bits_left < BITS_PER_SYMBOL {
					h_data = uint32(local_h[idx]) + (uint32(local_h[idx+1]) << BITS_PER_BYTE)
					idx += 2
					current_data += h_data << bits_left
					bits_left += 16
				}
				imag := current_data & 1023
				bits_left -= BITS_PER_SYMBOL
				current_data = current_data >> BITS_PER_SYMBOL

				if bits_left < BITS_PER_SYMBOL {
					h_data = uint32(local_h[idx]) + (uint32(local_h[idx+1]) << BITS_PER_BYTE)
					idx += 2
					current_data += h_data << bits_left
					bits_left += 16
				}

				real := current_data & 1023
				bits_left -= BITS_PER_SYMBOL
				current_data = current_data >> BITS_PER_SYMBOL

				csi[nr_idx+nc_idx*2][k] = complex(float64(bitConvert(int(real))), float64(bitConvert(int(imag))))
			}
		}
	}
	return csi
}
