package databuffer

import "csidealer/pkg/csicore"

type PacketBuffer struct {
	Data [][]csicore.CsiPackage
}

func NewPacketBuffer() *PacketBuffer {
	p := new(PacketBuffer)
	return p
}

func (buf *PacketBuffer) Push(data []csicore.CsiPackage) {
	buf.Data = append(buf.Data, data)
}

func (buf *PacketBuffer) Shift(n int) []csicore.CsiPackage {
	dataPart := buf.Data[0]
	buf.Data = buf.Data[1:]
	return dataPart
}

func (buf PacketBuffer) Length() int {
	return len(buf.Data)
}