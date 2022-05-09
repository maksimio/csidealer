package repo

import (
	"csidealer/internal/entity"
)

type CsiLocalRepo struct {
	data      []*entity.Package
	fullCount uint64
	maxCount  uint64
}

func NewCsiLocalRepo(maxCount uint64) *CsiLocalRepo {
	p := &CsiLocalRepo{
		maxCount: maxCount,
		data:     make([]*entity.Package, 0, maxCount),
	}
	return p
}

func (c *CsiLocalRepo) Push(csiPackage *entity.Package) {
	c.data = append(c.data, csiPackage)
	c.fullCount += 1
	if c.fullCount > c.maxCount {
		c.data = c.data[1:]
	}
}

func (c *CsiLocalRepo) GetLastN(n int) []*entity.Package {
	length := len(c.data)
	if n > length {
		n = length
	}

	return c.data[length-n:]
}

func (c *CsiLocalRepo) GetFullCount() uint64 {
	return c.fullCount
}
