package databuffer

import (
	"csidealer/pkg/csicore"
	"math/cmplx"
	"sync"
)

const MAX_COUNT = 20

type PackageBuffer struct {
	Data      []*Package
	c         <-chan csicore.CsiPackage
	fullCount uint64
	mutex     sync.Mutex
}

func NewPackageBuffer(c <-chan csicore.CsiPackage) *PackageBuffer {
	p := new(PackageBuffer)
	p.c = c
	return p
}

func (buf *PackageBuffer) Listen() {
	for onePackage := range buf.c {
		buf.push(onePackage)
	}
}

func (buf *PackageBuffer) push(data csicore.CsiPackage) {
	data.Abs = csicore.CsiMap(data.Csi, cmplx.Abs)
	data.Phase = csicore.CsiMap(data.Csi, cmplx.Phase)
	data.Re = csicore.CsiMap(data.Csi, realWrapper)
	data.Im = csicore.CsiMap(data.Csi, imagWrapper)

	buf.mutex.Lock()
	buf.Data = append(buf.Data, NewPackage(data, buf.fullCount))
	buf.fullCount += 1
	if buf.fullCount > MAX_COUNT {
		buf.Data = buf.Data[1:]
	}
	buf.mutex.Unlock()
}

func (buf *PackageBuffer) Length() int {
	return len(buf.Data)
}

func (buf *PackageBuffer) LastN(n int) {
	length := buf.Length()
	if n > length {
		n = length
	}

	// return buf.Data[length-n:]
}

func realWrapper(c complex128) float64 {
	return real(c)
}

func imagWrapper(c complex128) float64 {
	return imag(c)
}
