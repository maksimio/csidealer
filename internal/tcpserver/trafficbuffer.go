package tcpserver

type TrafficBuffer struct {
	Data []byte
	Size int
}


func (buf *TrafficBuffer) Push(data []byte) {
	buf.Data = append(buf.Data, data...)
}

func (buf *TrafficBuffer) Shift(n int) {
	// buf.Data = append(buf.Data, data...)
}