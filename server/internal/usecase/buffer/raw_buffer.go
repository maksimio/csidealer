package buffer

import (
	"encoding/binary"
)

const _sizeByteLen = 4

type CsiRawBuffer struct {
	rawData         []byte
	nextPackageSize int
	splittedData    [][]byte
}

func NewCsiRawRepo() *CsiRawBuffer {
	return &CsiRawBuffer{}
}

func (c *CsiRawBuffer) Push(data []byte) {
	c.rawData = append(c.rawData, data...)
	c.splitPackageAll()
}

func (c *CsiRawBuffer) GetAllSplitted() [][]byte {
	temp := c.splittedData
	c.splittedData = [][]byte{}
	return temp
}

func (c *CsiRawBuffer) splitPackageAll() {
	c.splitPackage()
	for c.nextPackageSize+_sizeByteLen < len(c.rawData) {
		c.splitPackage()
	}
}

func (c *CsiRawBuffer) splitPackage() {
	if c.nextPackageSize == 0 && len(c.rawData) >= _sizeByteLen {
		c.nextPackageSize = int(binary.LittleEndian.Uint32(c.shift(_sizeByteLen)))
	} else if c.nextPackageSize != 0 && c.nextPackageSize <= len(c.rawData) {
		data := c.shift(c.nextPackageSize)
		c.nextPackageSize = 0

		c.splittedData = append(c.splittedData, data)
	}
}

func (c *CsiRawBuffer) shift(n int) []byte {
	dataPart := c.rawData[:n]
	c.rawData = c.rawData[n:]
	return dataPart
}
