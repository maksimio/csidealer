package models

type Mark struct {
	Id                  string `json:"id"`
	Text                string `json:"text"`
	IsActive            bool   `json:"is_active"`
	Timestamp           int64  `json:"timestamp"`
	DeltaTimestamp      int64  `json:"delta_timestamp"`
	CsiPackageNum       uint64 `json:"csi_package_num"`
	CsiPackageTimestamp uint64 `json:"csi_package_timestamp"`
}
