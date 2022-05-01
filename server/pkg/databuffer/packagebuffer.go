package databuffer

import (
	"csidealer/pkg/csicore"
	"sync"
)

const MAX_COUNT = 20

type PackageBuffer struct {
	Data      []csicore.CsiPackage
	c         <-chan *csicore.CsiPackage
	fullCount uint64
	mutex     sync.Mutex
}

func NewPackageBuffer(c <-chan *csicore.CsiPackage) *PackageBuffer {
	p := new(PackageBuffer)
	p.c = c
	return p
}

func (buf *PackageBuffer) Listen() {
	for onePackage := range buf.c {
		buf.push(onePackage)
	}
}

func (buf *PackageBuffer) push(data *csicore.CsiPackage) {
	buf.mutex.Lock()
	buf.Data = append(buf.Data, *csicore.NewCsiPackage(1))
	buf.fullCount += 1
	if buf.fullCount > MAX_COUNT {
		buf.Data = buf.Data[1:]
	}
	buf.mutex.Unlock()
}

func (buf *PackageBuffer) Length() int {
	return len(buf.Data)
}

func (buf *PackageBuffer) LastN(n int) []csicore.CsiPackage {
	length := buf.Length()
	if n > length {
		n = length
	}

	return buf.Data[length-n:]
}
