package filter

import "csidealer/internal/models"

type FilterService struct {
	payloadLenMin uint16
	payloadLenMax uint16
	nr            uint8
	nc            uint8
	nTones        uint8
	in            <-chan models.Package
	outs          []chan<- models.Package
}

func NewFilterService(in <-chan models.Package, outs []chan<- models.Package, payloadLenMin, payloadLenMax uint16, nr, nc, nTones uint8) *FilterService {
	return &FilterService{
		payloadLenMin: payloadLenMin,
		payloadLenMax: payloadLenMax,
		nr:            nr,
		nc:            nc,
		nTones:        nTones,
		in:            in,
		outs:          outs,
	}
}

func (f *FilterService) Run() {
	for {
		pack := <-f.in

		if !f.allow(pack.Info) {
			continue // осознанный прием, чтобы избегать вложенности
		}

		for _, out := range f.outs {
			out <- pack
		}
	}
}

func (f *FilterService) allow(info *models.PackageInfo) bool {
	return info.PayloadLength >= f.payloadLenMin &&
		info.PayloadLength <= f.payloadLenMax &&
		info.Nr == f.nr &&
		info.Nc == f.nc &&
		info.NumTones == f.nTones
}

func (f *FilterService) GetLimits() (payloadLenMin, payloadLenMax uint16, nr, nc, nTones uint8) {
	return f.payloadLenMin, f.payloadLenMax, f.nr, f.nc, f.nTones
}

func (f *FilterService) SetLimits(payloadLenMin, payloadLenMax uint16, nr, nc, nTones uint8) {
	f.payloadLenMin = payloadLenMin
	f.payloadLenMax = payloadLenMax
	f.nr = nr
	f.nc = nc
	f.nTones = nTones
}
