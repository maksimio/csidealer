package databuffer

type PackageBuffer struct {
	Data [][]byte
	CurrentSize int
}

func NewPackageBuffer() *PackageBuffer {
	p := new(PackageBuffer)
	return p
}

func (buf *PackageBuffer) Push(data []byte) {
	buf.Data = append(buf.Data, data)
}

func (buf *PackageBuffer) Shift() []byte {
	dataPart := buf.Data[0]
	buf.Data = buf.Data[1:]
	return dataPart
}

func (buf PackageBuffer) Length() int {
	return len(buf.Data)
}