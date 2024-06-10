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
	outs               []chan<- models.RawPackage
}

func NewBufferService(outs []chan<- models.RawPackage) *BufferService {
	return &BufferService{
		outs: outs,
	}
}

func (c *BufferService) Push(data []byte) {
	c.rawData = append(c.rawData, data...)
	c.splitPackages()

	for _, d := range c.splittedData {
		for _, out := range c.outs {
			out <- d
		}
	}
	c.splittedData = []models.RawPackage{}
}

func (c *BufferService) Flush() {
	c.rawData = []byte{}
	c.currentPackageSize = 0
	c.splittedData = []models.RawPackage{}
}

func (c *BufferService) splitPackages() {
	for len(c.rawData) >= c.currentPackageSize && len(c.rawData) > _sizeByteLen {
		if c.currentPackageSize == 0 && len(c.rawData) >= _sizeByteLen {
			c.currentPackageSize = int(binary.LittleEndian.Uint32(c.shift(_sizeByteLen)))
		}

		if c.currentPackageSize != 0 && len(c.rawData) >= c.currentPackageSize {
			data := c.shift(c.currentPackageSize)

			c.splittedData = append(c.splittedData, models.RawPackage{
				/* TCP-клиент client_main на роутере Rx, который пересылает данные, по неизвестной причине
				в конце data дублирует еще 2 байта размера пакета. Программа recv_csi так не делает, это, скорее всего, ошибка.
				Библиотека csi_read опирается не на поле Size для управления смещением указателя в файле,
				а просто на последовательное чтение нужного числа байт. Из-за этого в ней будет ошибка,
				если здесь не обрезать data[:len(data)-2]. Также уменьшим Size, хотя это на csiread не влияет */
				Data: data[:len(data)-2],
				Size: uint16(c.currentPackageSize - 2),
			})
			c.currentPackageSize = 0
		}
	}
}

func (c *BufferService) shift(n int) []byte {
	dataPart := c.rawData[:n]
	c.rawData = c.rawData[n:]
	return dataPart
}
