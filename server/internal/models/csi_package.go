package models

type RawPackage struct {
	Size uint16
	Data []byte
}

type Csi [][]complex128

type Package struct {
	Timestamp int64
	Uuid      string
	Number    uint64

	Info *PackageInfo
	Data Csi
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

const (
	CSI_ABS uint8 = iota
	CSI_PHASE
	CSI_IM
	CSI_RE
	CSI_PHASE_WITHOUT_JUMPS
)

const (
	FEATURE_MEAN uint8 = iota
	FEATURE_MAX
	FEATURE_MIN
	FEATURE_MEDIAN
	FEATURE_STD
	FEATURE_SKEWNESS
	FEATURE_KURTOSIS
)

type ApiPackage struct {
	Timestamp int64        `json:"ts"`
	Id        string       `json:"id"`
	Number    uint64       `json:"n"`
	Info      *PackageInfo `json:"info"`
	Data      [][]float64  `json:"data"`
}

type ApiPackageAbsPhase struct {
	Timestamp int64        `json:"ts"`
	Id        string       `json:"id"`
	Number    uint64       `json:"n"`
	Info      *PackageInfo `json:"info"`
	Abs       [][]float64  `json:"abs"`
	Phase     [][]float64  `json:"phase"`
}
