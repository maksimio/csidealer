package buffer

import (
	"csidealer/internal/models"
	"encoding/binary"
)

const _sizeByteLen = 4

type BufferService struct {
	rawData            []byte
	currentPackageSize int
	splittedData       []models.RawPackage
	TcpRemoteAddr      string
	out                chan<- models.RawPackage
}

func NewBufferService(out chan<- models.RawPackage) *BufferService {
	return &BufferService{}
}

func (c *BufferService) Push(data []byte) {
	c.rawData = append(c.rawData, data...)
	c.splitPackageAll()

	for _, d := range c.splittedData {
		c.out <- d
	}
	c.splittedData = []models.RawPackage{}
}

func (c *BufferService) GetAllSplitted() []models.RawPackage {
	temp := c.splittedData
	c.splittedData = []models.RawPackage{}
	return temp
}

func (c *BufferService) Flush() {
	c.rawData = []byte{}
	c.currentPackageSize = 0
	c.splittedData = []models.RawPackage{}
}

func (c *BufferService) splitPackageAll() {
	c.splitPackage()
	for c.currentPackageSize+_sizeByteLen < len(c.rawData) {
		c.splitPackage()
	}
}

func (c *BufferService) splitPackage() {
	if c.currentPackageSize == 0 && len(c.rawData) >= _sizeByteLen {
		c.currentPackageSize = int(binary.LittleEndian.Uint32(c.shift(_sizeByteLen)))
	} else if c.currentPackageSize != 0 && c.currentPackageSize <= len(c.rawData) {
		data := c.shift(c.currentPackageSize)

		c.splittedData = append(c.splittedData, models.RawPackage{
			Data: data,
			Size: uint16(c.currentPackageSize),
		})
		c.currentPackageSize = 0
	}
}

func (c *BufferService) shift(n int) []byte {
	dataPart := c.rawData[:n]
	c.rawData = c.rawData[n:]
	return dataPart
}
