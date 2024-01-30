package decoder

import "csidealer/internal/models"

type DecoderService struct {
	in   <-chan models.RawPackage
	outs []chan<- models.Package
}

func NewDecoderService(in <-chan models.RawPackage, outs []chan<- models.Package) *DecoderService {
	return &DecoderService{
		in:   in,
		outs: outs,
	}
}

func (d *DecoderService) Run() {
	for {
		rawPackage := <-d.in
		pack := decodeCsiPackage(rawPackage.Data)
		for _, out := range d.outs {
			out <- pack
		}
	}
}

func decodeCsiPackage(data []byte) models.Package {
	pack := models.Package{
		Info: decodePackageInfo(data),
	}

	if pack.Info.CsiLength > 0 {
		rawCsi := data[SHIFT_CSI_INFO : SHIFT_CSI_INFO+pack.Info.CsiLength]
		pack.Data = decodeCsi(rawCsi, pack.Info.Nr, pack.Info.Nc, pack.Info.NumTones)
	}

	return pack
}
