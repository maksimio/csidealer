package databuffer

import (
	"csidealer/pkg/csi"
	"time"

	"github.com/google/uuid"
)

type Package struct {
	Timestamp int64
	Uuid      string
	Number    uint64
	CsiPack   csi.CsiPackage
}

func NewPackage(data csi.CsiPackage, number uint64) *Package {
	p := new(Package)
	p.CsiPack = data
	p.Number = number
	p.Uuid = uuid.New().String()
	p.Timestamp = time.Now().UnixMilli()
	return p
}
