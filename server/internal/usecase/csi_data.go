package usecase

import (
	"csidealer/internal/entity"
	"errors"
)

func (uc *CsiUseCase) GetCsiPackageCount() uint64 {
	return uc.repo.GetFullCount()
}

func (uc *CsiUseCase) GetCsi(csiType uint8, count int) ([]entity.ApiPackage, error) {
	packets := uc.repo.GetLastN(count)

	switch csiType {
	case entity.CSI_ABS:
		return uc.proc.Abs(packets), nil
	case entity.CSI_PHASE:
		return uc.proc.Phase(packets), nil
	case entity.CSI_IM:
		return uc.proc.Im(packets), nil
	case entity.CSI_RE:
		return uc.proc.Re(packets), nil
	default:
		return []entity.ApiPackage{}, errors.New("указан неверный тип данных")
	}
}
