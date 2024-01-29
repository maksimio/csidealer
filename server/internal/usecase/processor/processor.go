package processor

import (
	entity "csidealer/internal/models"
	"errors"
	"math"
)

type Processor struct {
	rounder float64
}

func NewProcessor(rounder int) *Processor {
	return &Processor{
		rounder: math.Pow10(rounder),
	}
}

func (p *Processor) CsiMap(csi entity.Csi, f func(complex128) float64) [][]float64 {
	data := make([][]float64, len(csi))
	for i := range csi {
		data[i] = make([]float64, len(csi[i]))

		for j, value := range csi[i] {
			data[i][j] = math.Round(f(value)*p.rounder) / p.rounder
		}
	}

	return data
}

func (p *Processor) PackageMap(data []*entity.Package, handler func(complex128) float64) []entity.ApiPackage {
	packs := make([]entity.ApiPackage, 0, len(data))

	for _, value := range data {
		packs = append(packs, entity.ApiPackage{
			Timestamp: value.Timestamp,
			Id:        value.Uuid,
			Info:      value.Info,
			Number:    value.Number,
			Data:      p.CsiMap(value.Data, handler),
		})
	}

	return packs
}

func (p *Processor) SubcarrierMap(data []*entity.Package, handler func(complex128) float64, h, i int) ([]float64, error) {
	subcarrierData := make([]float64, 0, len(data))

	for _, pack := range data {
		if h >= len(pack.Data) {
			return []float64{}, errors.New("h: выход за границы массива")
		}
		hData := pack.Data[h]
		if i >= len(hData) {
			return []float64{}, errors.New("i: выход за границы массива")
		}
		value := hData[i]
		subcarrierData = append(subcarrierData, handler(value))
	}

	return subcarrierData, nil
}
