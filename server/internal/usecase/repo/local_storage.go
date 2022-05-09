package repo

import (
	"csidealer/internal/entity"
)

type CsiLocalRepo struct {
	Data      []*entity.Package
	fullCount uint64
	maxCount  uint64
}

func NewCsiLocalRepo(maxCount uint64) *CsiLocalRepo {
	p := &CsiLocalRepo{
		maxCount: maxCount,
		Data:     make([]*entity.Package, maxCount),
	}
	return p
}

func (c *CsiLocalRepo) Push(csiPackage *entity.Package) {
	c.Data = append(c.Data, csiPackage)
	c.fullCount += 1
	if c.fullCount > c.maxCount {
		c.Data = c.Data[1:]
	}
}

func (c *CsiLocalRepo) GetLastN(n int) []*entity.Package {
	length := len(c.Data)
	if n > length {
		n = length
	}

	return c.Data[length-n:]
}
