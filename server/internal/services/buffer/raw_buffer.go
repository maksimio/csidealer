package buffer

import (
	entity "csidealer/internal/models"
	"encoding/binary"
)

const _sizeByteLen = 4

type CsiRawBuffer struct {
	rawData            []byte
	currentPackageSize int
	splittedData       []entity.RawPackage
}

func NewCsiRawRepo() *CsiRawBuffer {
	return &CsiRawBuffer{}
}

func (c *CsiRawBuffer) Push(data []byte) {
	c.rawData = append(c.rawData, data...)
	c.splitPackageAll()
}

func (c *CsiRawBuffer) GetAllSplitted() []entity.RawPackage {
	temp := c.splittedData
	c.splittedData = []entity.RawPackage{}
	return temp
}

func (c *CsiRawBuffer) Flush() {
	c.rawData = []byte{}
	c.currentPackageSize = 0
	c.splittedData = []entity.RawPackage{}
}

func (c *CsiRawBuffer) splitPackageAll() {
	c.splitPackage()
	for c.currentPackageSize+_sizeByteLen < len(c.rawData) {
		c.splitPackage()
	}
}

func (c *CsiRawBuffer) splitPackage() {
	if c.currentPackageSize == 0 && len(c.rawData) >= _sizeByteLen {
		c.currentPackageSize = int(binary.LittleEndian.Uint32(c.shift(_sizeByteLen)))
	} else if c.currentPackageSize != 0 && c.currentPackageSize <= len(c.rawData) {
		data := c.shift(c.currentPackageSize)

		c.splittedData = append(c.splittedData, entity.RawPackage{
			Data: data,
			Size: uint16(c.currentPackageSize),
		})
		c.currentPackageSize = 0
	}
}

func (c *CsiRawBuffer) shift(n int) []byte {
	dataPart := c.rawData[:n]
	c.rawData = c.rawData[n:]
	return dataPart
}
