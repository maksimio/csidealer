package databuffer

import "sync"

type TrafficBuffer struct {
	Data []byte
	mutex sync.Mutex
}

func NewTrafficBuffer() *TrafficBuffer {
	p := new(TrafficBuffer)
	return p
}

func (buf *TrafficBuffer) Push(data []byte) {
	buf.mutex.Lock()
	buf.Data = append(buf.Data, data...)
	buf.mutex.Unlock()
}

func (buf *TrafficBuffer) Shift(n int) []byte {
	dataPart := buf.Data[:n]
	buf.mutex.Lock()
	buf.Data = buf.Data[n:]
	buf.mutex.Unlock()
	return dataPart
}

func (buf *TrafficBuffer) Length() int {
	return len(buf.Data)
}