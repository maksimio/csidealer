package databuffer

import (
	"encoding/binary"
	"fmt"
)

type BufferFlow struct {
	trafficBuf TrafficBuffer
	packageBuf PackageBuffer
	parseBuf   ParseBuffer
}

func NewBufferFlow() *BufferFlow {
	p := new(BufferFlow)
	p.trafficBuf = *NewTrafficBuffer()
	p.packageBuf = *NewPackageBuffer()
	p.parseBuf = *NewParseBuffer()
	return p
}

func (buf *BufferFlow) Push(data []byte) {
	buf.trafficBuf.Push(data)
	buf.splitPacketAll()
}

func (buf *BufferFlow) Length() int {
	return buf.trafficBuf.Length()
}

func (buf *BufferFlow) splitPacketAll() {
	buf.splitPacket()
	fmt.Println("CurSize:", buf.packageBuf.CurrentSize)
	for buf.packageBuf.CurrentSize+4 < uint32(buf.trafficBuf.Length()) {
		buf.splitPacket()
	}
}

func (buf *BufferFlow) splitPacket() {
	if buf.packageBuf.CurrentSize == 0 && buf.trafficBuf.Length() >= 4 {
		buf.packageBuf.CurrentSize = binary.LittleEndian.Uint32(buf.trafficBuf.Shift(4))
	} else if buf.packageBuf.CurrentSize != 0 && buf.packageBuf.CurrentSize <= uint32(buf.trafficBuf.Length()) {
		buf.packageBuf.Push(buf.trafficBuf.Shift(int(buf.packageBuf.CurrentSize)))
		fmt.Println(buf.packageBuf.CurrentSize)
		buf.packageBuf.CurrentSize = 0
	}
}
