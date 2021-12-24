package databuffer

type ParseBuffer struct {
	Data [][]byte
}

func NewParseBuffer() *ParseBuffer {
	p := new(ParseBuffer)
	return p
}

func (buf *ParseBuffer) Push(data []byte) {
	buf.Data = append(buf.Data, data)
}

func (buf *ParseBuffer) Shift(n int) []byte {
	dataPart := buf.Data[0]
	buf.Data = buf.Data[1:]
	return dataPart
}

func (buf ParseBuffer) Length() int {
	return len(buf.Data)
}