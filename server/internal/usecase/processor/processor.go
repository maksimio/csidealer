package processor

import (
	"csidealer/internal/entity"
	"math"
	"math/cmplx"
)

type Processor struct {
	rounder float64
}

func NewProcessor(rounder int) *Processor {
	return &Processor{
		rounder: math.Pow10(rounder),
	}
}

func (p *Processor) csiMap(csi entity.Csi, f func(complex128) float64) [][]float64 {
	data := make([][]float64, len(csi))
	for i := range csi {
		data[i] = make([]float64, len(csi[i]))

		for j, value := range csi[i] {
			data[i][j] = math.Round(f(value)*p.rounder) / p.rounder
		}
	}

	return data
}

func (p *Processor) Abs(data []*entity.Package) []entity.ApiPackage {
	packs := make([]entity.ApiPackage, 0, len(data))

	for _, value := range data {
		packs = append(packs, entity.ApiPackage{
			Timestamp: value.Timestamp,
			Id:        value.Uuid,
			Info:      value.Info,
			Number:    value.Number,
			Data:      p.csiMap(value.Data, cmplx.Abs),
		})
	}

	return packs
}

func (p *Processor) Phase(data []*entity.Package) []entity.ApiPackage {
	packs := make([]entity.ApiPackage, 0, len(data))

	for _, value := range data {
		packs = append(packs, entity.ApiPackage{
			Timestamp: value.Timestamp,
			Id:        value.Uuid,
			Info:      value.Info,
			Number:    value.Number,
			Data:      p.csiMap(value.Data, cmplx.Phase),
		})
	}

	return packs
}

func (p *Processor) Re(data []*entity.Package) []entity.ApiPackage {
	packs := make([]entity.ApiPackage, 0, len(data))

	for _, value := range data {
		packs = append(packs, entity.ApiPackage{
			Timestamp: value.Timestamp,
			Id:        value.Uuid,
			Info:      value.Info,
			Number:    value.Number,
			Data:      p.csiMap(value.Data, reWrapper),
		})
	}

	return packs
}

func (p *Processor) Im(data []*entity.Package) []entity.ApiPackage {
	packs := make([]entity.ApiPackage, 0, len(data))

	for _, value := range data {
		packs = append(packs, entity.ApiPackage{
			Timestamp: value.Timestamp,
			Id:        value.Uuid,
			Info:      value.Info,
			Number:    value.Number,
			Data:      p.csiMap(value.Data, imWrapper),
		})
	}

	return packs
}

func (p *Processor) PhaseWithoutJumps(data []*entity.Package) []entity.ApiPackage {
	return []entity.ApiPackage{}
}
