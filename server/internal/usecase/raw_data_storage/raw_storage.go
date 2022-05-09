package raw_data_storage

import (
	"encoding/binary"
)

type CsiRawRepo struct {
	rawData         []byte
	nextPackageSize int
	splittedData    [][]byte
}

func (c *CsiRawRepo) Push(data []byte) {
	c.rawData = append(c.rawData, data...)
	c.splitPackageAll()
}

func (c *CsiRawRepo) GetAll() [][]byte {
	temp := c.splittedData
	c.splittedData = [][]byte{}
	return temp
}

func (c *CsiRawRepo) splitPackageAll() {
	c.splitPackage()
	for c.nextPackageSize+4 < len(c.rawData) {
		c.splitPackage()
	}
}

func (c *CsiRawRepo) splitPackage() {
	if c.nextPackageSize == 0 && len(c.rawData) >= 4 {
		c.nextPackageSize = int(binary.LittleEndian.Uint32(c.shift(4)))
	} else if c.nextPackageSize != 0 && c.nextPackageSize <= len(c.rawData) {
		data := c.shift(c.nextPackageSize)
		c.nextPackageSize = 0

		c.splittedData = append(c.splittedData, data)
	}
}

func (c *CsiRawRepo) shift(n int) []byte {
	dataPart := c.rawData[:n]
	c.rawData = c.rawData[n:]
	return dataPart
}
