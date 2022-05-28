package processor

import "csidealer/internal/entity"

type Processor struct{}

func NewProcessor() *Processor {
	return &Processor{}
}

func (p *Processor) Abs(data []*entity.Package) []entity.ApiPackage {
	return []entity.ApiPackage{}
}

func (p *Processor) Phase(data []*entity.Package) []entity.ApiPackage {
	return []entity.ApiPackage{}
}

func (p *Processor) Re(data []*entity.Package) []entity.ApiPackage {
	return []entity.ApiPackage{}
}

func (p *Processor) Im(data []*entity.Package) []entity.ApiPackage {
	return []entity.ApiPackage{}
}

func (p *Processor) PhaseWithoutJumps(data []*entity.Package) []entity.ApiPackage {
	return []entity.ApiPackage{}
}
