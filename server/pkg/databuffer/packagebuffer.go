package databuffer

import (
	"csidealer/pkg/csicore"
	"fmt"
	"sync"
)

const MAX_COUNT = 20

type PackageBuffer struct {
	Data      []csicore.CsiPackage
	fullCount uint64
	mutex     sync.Mutex
}

func NewPackageBuffer() *PackageBuffer {
	p := new(PackageBuffer)
	return p
}

func (buf *PackageBuffer) Push(data csicore.CsiPackage) {
	buf.mutex.Lock()
	buf.Data = append(buf.Data, data)
	buf.fullCount += 1
	if buf.fullCount > MAX_COUNT {
		buf.Data = buf.Data[1:]
	}
	buf.mutex.Unlock()
	fmt.Println("CC:", buf.Length(), buf.fullCount)
}

func (buf *PackageBuffer) Length() int {
	return len(buf.Data)
}
