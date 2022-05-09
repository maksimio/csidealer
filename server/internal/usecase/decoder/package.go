package decoder

import "csidealer/internal/entity"

func DecodeCsiPackage(data []byte) entity.Package {
	var pack entity.Package
	pack.Info = decodePackageInfo(data)

	if pack.Info.CsiLength > 0 {
		rawCsi := data[SHIFT_CSI_INFO : SHIFT_CSI_INFO+pack.Info.CsiLength]
		pack.Data = decodeCsi(rawCsi, pack.Info.Nr, pack.Info.Nc, pack.Info.NumTones)
	}

	return pack
}
