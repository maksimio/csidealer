package csicore

import (
	"github.com/google/uuid"
	"time"
)

type CsiPackage struct {
	PackageInfo PackageInfo
	Csi         [][]complex128
	Abs         [][]float64
	Phase       [][]float64
	Timestamp int64
	Uuid      string
	Number    uint64
}

func NewCsiPackage(number uint64) *CsiPackage {
	p := new(CsiPackage)
	p.Number = number
	p.Uuid = uuid.New().String()
	p.Timestamp = time.Now().UnixMilli()
	return p
}

type PackageInfo struct {
	Timestamp     uint64
	CsiLength     uint16
	TxChannel     uint16
	ErrInfo       uint8
	NoiseFloor    uint8
	Rate          uint8
	BandWidth     uint8
	NumTones      uint8
	Nr            uint8
	Nc            uint8
	Rssi0         uint8
	Rssi1         uint8
	Rssi2         uint8
	Rssi3         uint8
	PayloadLength uint16
}
