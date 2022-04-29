package databuffer

import (
	"csidealer/pkg/csicore"
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
	for buf.packageBuf.CurrentSize+4 < buf.trafficBuf.Length() {
		buf.splitPacket()
	}
}

func (buf *BufferFlow) splitPacket() {
	if buf.packageBuf.CurrentSize == 0 && buf.trafficBuf.Length() >= 4 {
		buf.packageBuf.CurrentSize = int(binary.LittleEndian.Uint32(buf.trafficBuf.Shift(4)))
	} else if buf.packageBuf.CurrentSize != 0 && buf.packageBuf.CurrentSize <= buf.trafficBuf.Length() {
		buf.packageBuf.Push(buf.trafficBuf.Shift(buf.packageBuf.CurrentSize))
		buf.packageBuf.CurrentSize = 0
		buf.parsePacket()
	}
}

func (buf *BufferFlow) parsePacket() {
	data := buf.packageBuf.Shift()
	pack := csicore.DecodeCsiPackage(data)
	pack.Abs = csicore.CsiToAbs(pack.Csi)
	pack.Phase = csicore.CsiToPhase(pack.Csi)

	if pack.PackageInfo.CsiLength > 0 {
		fmt.Println(pack.Abs[0][20:30])
	}
}