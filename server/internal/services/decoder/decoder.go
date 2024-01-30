package decoder

import (
	"csidealer/internal/models"
	"time"

	"github.com/google/uuid"
)

type DecoderService struct {
	in               <-chan models.RawPackage
	outs             []chan<- models.Package
	csiPackageNumber uint64
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
		pack := d.decodeCsiPackage(rawPackage.Data)
		if pack.Info.CsiLength == 0 {
			continue
		}

		for _, out := range d.outs {
			out <- pack
		}
	}
}

func (d *DecoderService) decodeCsiPackage(data []byte) models.Package {
	info := decodePackageInfo(data)

	// if info.CsiLength == 0 {
	// } // TODO: не будет ли бага в decodeCsi?

	rawCsi := data[SHIFT_CSI_INFO : SHIFT_CSI_INFO+info.CsiLength]

	pack := models.Package{
		Info:      info,
		Data:      decodeCsi(rawCsi, info.Nr, info.Nc, info.NumTones),
		Uuid:      uuid.New().String(),
		Timestamp: time.Now().UnixMilli(),
		Number:    d.csiPackageNumber,
	}

	d.csiPackageNumber += 1
	return pack
}
