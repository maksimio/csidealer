package databuffer

type BufferFlow struct {
	trafficBuf TrafficBuffer
	parseBuf   ParseBuffer
	packageBuf PackageBuffer
}

func NewBufferFlow() *BufferFlow {
	p := new(BufferFlow)
	p.trafficBuf = *NewTrafficBuffer()
	p.parseBuf = *NewParseBuffer()
	p.packageBuf = *NewPackageBuffer()
	return p
}

func (buf *BufferFlow) Push(data []byte) {
	buf.trafficBuf.Push(data)
}

func (buf *BufferFlow) Length() int {
	return buf.trafficBuf.Length()
}