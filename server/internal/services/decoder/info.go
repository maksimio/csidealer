package decoder

import (
	entity "csidealer/internal/models"
	"encoding/binary"
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

func decodePackageInfo(data []byte) *entity.PackageInfo {
	return &entity.PackageInfo{
		Timestamp:     binary.BigEndian.Uint64(data),
		CsiLength:     binary.BigEndian.Uint16(data[SHIFT_CsiLength : SHIFT_CsiLength+2]),
		TxChannel:     binary.BigEndian.Uint16(data[SHIFT_TxChannel : SHIFT_TxChannel+2]),
		ErrInfo:       uint8(data[SHIFT_ErrInfo]),
		NoiseFloor:    uint8(data[SHIFT_NoiseFloor]),
		Rate:          uint8(data[SHIFT_Rate]),
		BandWidth:     uint8(data[SHIFT_BandWidth]),
		NumTones:      uint8(data[SHIFT_NumTones]),
		Nr:            uint8(data[SHIFT_Nr]),
		Nc:            uint8(data[SHIFT_Nc]),
		Rssi0:         uint8(data[SHIFT_Rssi0]),
		Rssi1:         uint8(data[SHIFT_Rssi1]),
		Rssi2:         uint8(data[SHIFT_Rssi2]),
		Rssi3:         uint8(data[SHIFT_Rssi3]),
		PayloadLength: binary.BigEndian.Uint16(data[SHIFT_Payloadlength : SHIFT_Payloadlength+2]),
	}
}
