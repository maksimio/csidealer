package databuffer

import (
	"csidealer/pkg/csicore"
	"github.com/google/uuid"
	"time"
)

type Package struct {
	Timestamp int64
	Uuid      string
	Number    uint64
	CsiPack   csicore.CsiPackage
}

func NewPackage(data csicore.CsiPackage, number uint64) *Package {
	p := new(Package)
	p.CsiPack = data
	p.Number = number
	p.Uuid = uuid.New().String()
	p.Timestamp = time.Now().UnixMilli()
	return p
}
