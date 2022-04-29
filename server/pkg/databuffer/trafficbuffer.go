package databuffer

type TrafficBuffer struct {
	Data []byte
}

func NewTrafficBuffer() *TrafficBuffer {
	p := new(TrafficBuffer)
	return p
}

func (buf *TrafficBuffer) Push(data []byte) {
	buf.Data = append(buf.Data, data...)
}

func (buf *TrafficBuffer) Shift(n int) []byte {
	dataPart := buf.Data[:n]
	buf.Data = buf.Data[n:]
	return dataPart
}

func (buf *TrafficBuffer) Length() int {
	return len(buf.Data)
}