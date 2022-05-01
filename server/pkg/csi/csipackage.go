package csi

type CsiPackage struct {
	PackageInfo PackageInfo
	Csi         [][]complex128
}

type PackageInfo struct {
	Timestamp     uint64 `json:"ts"`
	CsiLength     uint16 `json:"csilen"`
	TxChannel     uint16 `json:"txchan"`
	ErrInfo       uint8  `json:"err"`
	NoiseFloor    uint8  `json:"noise"`
	Rate          uint8  `json:"rate"`
	BandWidth     uint8  `json:"bwidth"`
	NumTones      uint8  `json:"ntones"`
	Nr            uint8  `json:"nr"`
	Nc            uint8  `json:"nc"`
	Rssi0         uint8  `json:"rssi0"`
	Rssi1         uint8  `json:"rssi1"`
	Rssi2         uint8  `json:"rssi2"`
	Rssi3         uint8  `json:"rssi3"`
	PayloadLength uint16 `json:"payloadlen"`
}
