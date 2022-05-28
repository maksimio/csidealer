package filter

import "csidealer/internal/entity"

type Filter struct {
	payloadLenMin uint16
	payloadLenMax uint16
	nr            uint8
	nc            uint8
	nTones        uint8
}

func NewFilter(payloadLenMin, payloadLenMax uint16, nr, nc, nTones uint8) *Filter {
	return &Filter{
		payloadLenMin: payloadLenMin,
		payloadLenMax: payloadLenMax,
		nr:            nr,
		nc:            nc,
		nTones:        nTones,
	}
}

func (f *Filter) Check(info *entity.PackageInfo) bool {
	return info.PayloadLength >= f.payloadLenMin &&
		info.PayloadLength <= f.payloadLenMax &&
		info.Nr == f.nr &&
		info.Nc == f.nc &&
		info.NumTones == f.nTones
}

func (f *Filter) GetLimits() (payloadLenMin, payloadLenMax uint16, nr, nc, nTones uint8) {
	return f.payloadLenMin, f.payloadLenMax, f.nr, f.nc, f.nTones
}

func (f *Filter) SetLimits(payloadLenMin, payloadLenMax uint16, nr, nc, nTones uint8) {
	f.payloadLenMin = payloadLenMin
	f.payloadLenMax = payloadLenMax
	f.nr = nr
	f.nc = nc
	f.nTones = nTones
}
