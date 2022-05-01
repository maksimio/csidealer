package csicore

import (
	"encoding/binary"
)

const (
	BITS_PER_BYTE   = 8
	BITS_PER_SYMBOL = 10
)

const (
	SHIFT_CsiLength     = 8
	SHIFT_TxChannel     = 10
	SHIFT_ErrInfo       = 12
	SHIFT_NoiseFloor    = 13
	SHIFT_Rate          = 14
	SHIFT_BandWidth     = 15
	SHIFT_NumTones      = 16
	SHIFT_Nr            = 17
	SHIFT_Nc            = 18
	SHIFT_Rssi0         = 19
	SHIFT_Rssi1         = 20
	SHIFT_Rssi2         = 21
	SHIFT_Rssi3         = 22
	SHIFT_Payloadlength = 23
	SHIFT_CSI_INFO      = 25
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

func DecodePackageInfo(data []byte) PackageInfo {
	var info PackageInfo

	info.Timestamp = binary.BigEndian.Uint64(data)
	info.CsiLength = binary.BigEndian.Uint16(data[SHIFT_CsiLength:])
	info.TxChannel = binary.BigEndian.Uint16(data[SHIFT_TxChannel:])
	info.ErrInfo = uint8(data[SHIFT_ErrInfo])
	info.NoiseFloor = uint8(data[SHIFT_NoiseFloor])
	info.Rate = uint8(data[SHIFT_Rate])
	info.BandWidth = uint8(data[SHIFT_BandWidth])
	info.NumTones = uint8(data[SHIFT_NumTones])
	info.Nr = uint8(data[SHIFT_Nr])
	info.Nc = uint8(data[SHIFT_Nc])
	info.Rssi0 = uint8(data[SHIFT_Rssi0])
	info.Rssi1 = uint8(data[SHIFT_Rssi1])
	info.Rssi2 = uint8(data[SHIFT_Rssi2])
	info.Rssi3 = uint8(data[SHIFT_Rssi3])
	info.PayloadLength = binary.BigEndian.Uint16(data[SHIFT_Payloadlength:])
	return info
}

func DecodeCsiPackage(data []byte) *CsiPackage {
	pack := NewCsiPackage(1)
	pack.PackageInfo = DecodePackageInfo(data)

	if pack.PackageInfo.CsiLength > 0 {
		rawCsi := data[SHIFT_CSI_INFO : SHIFT_CSI_INFO+pack.PackageInfo.CsiLength]
		pack.Csi = DecodeCsi(rawCsi, pack.PackageInfo.Nr, pack.PackageInfo.Nc, pack.PackageInfo.NumTones)
	}

	return pack
}
