package entity

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
	Data      [][]float64     `json:"data"`
}
