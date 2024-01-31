package decoder

import (
	"csidealer/internal/models"
	"errors"
	"log"
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
		pack, err := d.decodeCsiPackage(rawPackage.Data)
		if err != nil {
			log.Print(err)
			continue
		}

		for _, out := range d.outs {
			out <- pack
		}
	}
}

func (d *DecoderService) decodeCsiPackage(data []byte) (models.Package, error) {
	info := decodePackageInfo(data)

	if info.CsiLength == 0 {
		return models.Package{}, errors.New("нет данных CSI в пакете")
	}

	rawCsi := data[SHIFT_CSI_INFO : SHIFT_CSI_INFO+info.CsiLength]

	pack := models.Package{
		Info:      info,
		Data:      decodeCsi(rawCsi, info.Nr, info.Nc, info.NumTones),
		Uuid:      uuid.New().String(),
		Timestamp: time.Now().UnixMilli(),
		Number:    d.csiPackageNumber,
	}

	d.csiPackageNumber += 1
	return pack, nil
}
