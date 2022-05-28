package usecase

import (
	"csidealer/internal/entity"
	"csidealer/internal/usecase/processor"
	"errors"
)

func (uc *CsiUseCase) GetCsiPackageCount() uint64 {
	return uc.repo.GetFullCount()
}

func (uc *CsiUseCase) GetCsiPackageMaxCount() uint64 {
	return uc.repo.GetMaxCount()
}

func (uc *CsiUseCase) GetCsi(csiType uint8, count int) ([]entity.ApiPackage, error) {
	packets := uc.repo.GetLastN(count)

	switch csiType {
	case entity.CSI_ABS:
		return uc.proc.PackageMap(packets, processor.AbsHandler), nil
	case entity.CSI_PHASE:
		return uc.proc.PackageMap(packets, processor.PhaseHandler), nil
	case entity.CSI_IM:
		return uc.proc.PackageMap(packets, processor.ImHandler), nil
	case entity.CSI_RE:
		return uc.proc.PackageMap(packets, processor.ReHandler), nil
	default:
		return []entity.ApiPackage{}, errors.New("указан неверный тип данных")
	}
}

func (uc *CsiUseCase) GetSubcarrier(csiType uint8, count, h, i int) ([]float64, error) {
	packets := uc.repo.GetLastN(count)

	switch csiType {
	case entity.CSI_ABS:
		return uc.proc.SubcarrierMap(packets, processor.AbsHandler, h, i)
	case entity.CSI_PHASE:
		return uc.proc.SubcarrierMap(packets, processor.PhaseHandler, h, i)
	case entity.CSI_IM:
		return uc.proc.SubcarrierMap(packets, processor.ImHandler, h, i)
	case entity.CSI_RE:
		return uc.proc.SubcarrierMap(packets, processor.ReHandler, h, i)
	default:
		return []float64{}, errors.New("указан неверный тип данных")
	}
}
