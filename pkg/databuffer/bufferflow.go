package databuffer

import (
	"csidealer/pkg/csicore"
	"encoding/binary"
	"fmt"
)

type BufferFlow struct {
	trafficBuf TrafficBuffer
	packetBuf   PacketBuffer
	nextPacketSize int
}

func NewBufferFlow() *BufferFlow {
	p := new(BufferFlow)
	p.trafficBuf = *NewTrafficBuffer()
	p.packetBuf = *NewPacketBuffer()
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
	for buf.nextPacketSize+4 < buf.trafficBuf.Length() {
		buf.splitPacket()
	}
}

func (buf *BufferFlow) splitPacket() {
	if buf.nextPacketSize == 0 && buf.trafficBuf.Length() >= 4 {
		buf.nextPacketSize = int(binary.LittleEndian.Uint32(buf.trafficBuf.Shift(4)))
	} else if buf.nextPacketSize != 0 && buf.nextPacketSize <= buf.trafficBuf.Length() {
		data := buf.trafficBuf.Shift(buf.nextPacketSize)
		buf.nextPacketSize = 0

		pack := csicore.DecodeCsiPackage(data)
		pack.Abs = csicore.CsiToAbs(pack.Csi)
		pack.Phase = csicore.CsiToPhase(pack.Csi)
	
		if pack.PackageInfo.CsiLength > 0 {
			fmt.Println(pack.Abs[0][20:30])
		}
	}
}