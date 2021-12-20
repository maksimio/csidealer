package databuffer

type PackageBuffer struct {
	Data [][]byte
}

func NewPackageBuffer() *PackageBuffer {
	p := new(PackageBuffer)
	return p
}

func (buf *PackageBuffer) Push(data []byte) {
	buf.Data = append(buf.Data, data)
}

func (buf *PackageBuffer) Shift(n int) [][]byte {
	dataPart := buf.Data[:1]
	buf.Data = buf.Data[1:]
	return dataPart
}

func (buf PackageBuffer) Length() int {
	return len(buf.Data)
}