package csicore

import (
	"encoding/binary"
	"fmt"
)

const (
	BITS_PER_BYTE   = 8
	BITS_PER_SYMBOL = 10
	SHIFT_CSI_INFO  = 25
)

func bitConvert(data int) int {
	if data&512 != 0 {
		data -= 1024
	}
	return data
}

func DecodeCsi(dataCsi []byte, nr, nc, numTones uint8) [][]complex128 {
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
		for ; ncIdx < nc; ncIdx++ {
			for ; nrIdx < nc; nrIdx++ {
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

func DecodePackageInfo(data []byte) PackageInfo {
	var info PackageInfo
	var shift uint8

	info.Timestamp = binary.BigEndian.Uint64(data)
	shift += 8
	info.CsiLength = binary.BigEndian.Uint16(data[shift:])
	shift += 2
	info.TxChannel = binary.BigEndian.Uint16(data[shift:])
	shift += 2
	info.ErrInfo = uint8(data[shift])
	shift += 1
	info.NoiseFloor = uint8(data[shift])
	shift += 1
	info.Rate = uint8(data[shift])
	shift += 1
	info.BandWidth = uint8(data[shift])
	shift += 1
	info.NumTones = uint8(data[shift])
	shift += 1
	info.Nr = uint8(data[shift])
	shift += 1
	info.Nc = uint8(data[shift])
	shift += 1
	info.Rssi0 = uint8(data[shift])
	shift += 1
	info.Rssi1 = uint8(data[shift])
	shift += 1
	info.Rssi2 = uint8(data[shift])
	shift += 1
	info.Rssi3 = uint8(data[shift])
	shift += 1
	info.Payloadlength = binary.BigEndian.Uint16(data[shift:])
	shift += 2

	fmt.Println("SHIFT:", shift)
	fmt.Println("CSILEN:", info.CsiLength)
	fmt.Println("PAYLOADLEN:", info.Payloadlength)
	fmt.Println(info)

	return info
}

func DecodeCsiPackage(data []byte) CsiPackage {
	var pack CsiPackage
	pack.PackageInfo = DecodePackageInfo(data)

	rawCsi := data[SHIFT_CSI_INFO : SHIFT_CSI_INFO+pack.PackageInfo.CsiLength]
	if pack.PackageInfo.CsiLength > 0 {
		pack.Csi = DecodeCsi(rawCsi, pack.PackageInfo.Nr, pack.PackageInfo.Nc, pack.PackageInfo.NumTones)
		fmt.Println("decoded CSI!")
	}

	fmt.Println(pack.Csi)
	return pack
}
