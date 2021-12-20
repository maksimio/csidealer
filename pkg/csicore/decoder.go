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

func DecodeCsi(data []byte, nr, nc, numTones uint8) {
	csi := make([][]complex128, nr*nc)
	for i := range csi {
		csi[i] = make([]complex128, numTones)
	}

	var bitsLeft uint8 = 16

	hData := data[0] + (data[1] << BITS_PER_BYTE)
	currentData := hData & 65535
	idx := 2

	var k, ncIdx, nrIdx uint8 = 0, 0, 0
	for ; k < numTones; k++ {
		for ; ncIdx < nc; ncIdx++ {
			for ; nrIdx < nc; nrIdx++ {
				
			}
		}
	}

}
