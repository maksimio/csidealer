package databuffer

import (
	"csidealer/pkg/csicore"
	"encoding/binary"
	"fmt"
	"math/cmplx"
)

type BufferFlow struct {
	trafficBuf      TrafficBuffer
	packageBuf      PackageBuffer
	nextPackageSize int
}

func NewBufferFlow() *BufferFlow {
	p := new(BufferFlow)
	p.trafficBuf = *NewTrafficBuffer()
	p.packageBuf = *NewPackageBuffer()
	return p
}

func (buf *BufferFlow) Push(data []byte) {
	buf.trafficBuf.Push(data)
	buf.splitPackageAll()
}

func (buf *BufferFlow) Length() int {
	return buf.trafficBuf.Length()
}

func (buf *BufferFlow) splitPackageAll() {
	buf.splitPackage()
	for buf.nextPackageSize+4 < buf.trafficBuf.Length() {
		buf.splitPackage()
	}
}

func (buf *BufferFlow) splitPackage() {
	if buf.nextPackageSize == 0 && buf.trafficBuf.Length() >= 4 {
		buf.nextPackageSize = int(binary.LittleEndian.Uint32(buf.trafficBuf.Shift(4)))
	} else if buf.nextPackageSize != 0 && buf.nextPackageSize <= buf.trafficBuf.Length() {
		data := buf.trafficBuf.Shift(buf.nextPackageSize)
		buf.nextPackageSize = 0
		go buf.decode(data)
	}
}

func (buf *BufferFlow)decode(data []byte) {
	pack := csicore.DecodeCsiPackage(data)
	pack.Abs = csicore.CsiMap(pack.Csi, cmplx.Abs)
	pack.Phase = csicore.CsiMap(pack.Csi, cmplx.Phase)

	buf.packageBuf.Push(pack)

	if pack.PackageInfo.CsiLength > 0 {
		fmt.Println(pack.Abs[0][20:30])
	}
}
