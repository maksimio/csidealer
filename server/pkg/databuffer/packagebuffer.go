package databuffer

import (
	"csidealer/pkg/csi"
	"csidealer/pkg/datatype"
	"sync"
)

const MAX_COUNT = 20

type PackageBuffer struct {
	Data      []*PackageUnion
	c         <-chan csi.CsiPackage
	fullCount uint64
	mutex     sync.Mutex
}

func NewPackageBuffer(c <-chan csi.CsiPackage) *PackageBuffer {
	p := new(PackageBuffer)
	p.c = c
	return p
}

func (buf *PackageBuffer) Listen() {
	for onePackage := range buf.c {
		buf.push(onePackage)
	}
}

func (buf *PackageBuffer) push(data csi.CsiPackage) {
	buf.mutex.Lock()
	buf.Data = append(buf.Data, NewPackageUnion(data, buf.fullCount))
	buf.fullCount += 1
	if buf.fullCount > MAX_COUNT {
		buf.Data = buf.Data[1:]
	}
	buf.mutex.Unlock()
}

func (buf *PackageBuffer) Length() int {
	return len(buf.Data)
}

func (buf *PackageBuffer) CsiLastN(n int, csiType string) []Package {
	length := buf.Length()
	if n > length {
		n = length
	}

	var packages []Package

	for i := length - n; i < length; i++ {
		var data Package

		switch csiType {
		case dataType.CsiDataType.Abs:
			data = buf.Data[i].Abs
		case dataType.CsiDataType.Phase:
			data = buf.Data[i].Phase
		case dataType.CsiDataType.Re:
			data = buf.Data[i].Re
		case dataType.CsiDataType.Im:
			data = buf.Data[i].Im
		}

		packages = append(packages, data)
	}

	return packages
}

func (buf *PackageBuffer) SubcarrierLastN(h, index, n int, csiType string) []float64 {
	length := buf.Length()
	if n > length {
		n = length
	}

	var data []float64

	for i := length - n; i < length; i++ {
		var value float64

		switch csiType {
		case dataType.CsiDataType.Abs:
			value = buf.Data[i].Abs.Data[h][index]
		case dataType.CsiDataType.Phase:
			value = buf.Data[i].Phase.Data[h][index]
		case dataType.CsiDataType.Re:
			value = buf.Data[i].Re.Data[h][index]
		case dataType.CsiDataType.Im:
			value = buf.Data[i].Im.Data[h][index]
		}

		data = append(data, value)
	}

	return data
}
