package databuffer

import (
	"csidealer/pkg/csicore"
	"time"
	"github.com/google/uuid"
)
type Package struct {
	timestamp int64
	uuid      string
	number    uint64
	Data      csicore.CsiPackage
}

func NewPackage(data csicore.CsiPackage, number uint64) *Package {
	p := new(Package)
	p.Data = data
	p.number = number
	p.uuid = uuid.New().String()
	p.timestamp = time.Now().UnixMilli()
	return p
}
