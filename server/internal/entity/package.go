package entity

type Package struct {
	Timestamp int64
	Uuid      string
	Number    uint64

	Info PackageInfo
	Data Csi
}
