package databuffer

import (
	"csidealer/pkg/csicore"
	"encoding/binary"
)

type TrafficBuffer struct {
	Data            []byte
	c               chan<- csicore.CsiPackage
	nextPackageSize int
}

func NewBufferFlow(c chan<- csicore.CsiPackage) *TrafficBuffer {
	p := new(TrafficBuffer)
	p.c = c
	return p
}

func (buf *TrafficBuffer) Push(data []byte) {
	buf.Data = append(buf.Data, data...)
	buf.splitPackageAll()
}

func (buf *TrafficBuffer) Length() int {
	return len(buf.Data)
}

func (buf *TrafficBuffer) splitPackageAll() {
	buf.splitPackage()
	for buf.nextPackageSize+4 < buf.Length() {
		buf.splitPackage()
	}
}

func (buf *TrafficBuffer) splitPackage() {
	if buf.nextPackageSize == 0 && buf.Length() >= 4 {
		buf.nextPackageSize = int(binary.LittleEndian.Uint32(buf.shift(4)))
	} else if buf.nextPackageSize != 0 && buf.nextPackageSize <= buf.Length() {
		data := buf.shift(buf.nextPackageSize)
		buf.nextPackageSize = 0
		go buf.decode(data)
	}
}

func (buf *TrafficBuffer) decode(data []byte) {
	pack := csicore.DecodeCsiPackage(data)
	if pack.PackageInfo.CsiLength == 0 {
		return
	}

	buf.c <- pack
}

func (buf *TrafficBuffer) shift(n int) []byte {
	dataPart := buf.Data[:n]
	buf.Data = buf.Data[n:]
	return dataPart
}
