package usecase

import (
	"csidealer/internal/entity"
	"errors"
)

func (uc *CsiUseCase) GetCsi(csiType uint8, count int) ([]entity.ApiPackage, error) {
	packets := uc.repo.GetLastN(count)

	switch csiType {
	case entity.CSI_ABS:
		return uc.proc.Abs(packets), nil
	case entity.CSI_PHASE:
		break
	case entity.CSI_IM:
		break
	case entity.CSI_RE:
		break
	default:
		return []entity.ApiPackage{}, errors.New("указан неверный тип данных")
	}

	return []entity.ApiPackage{}, errors.New("неизвестная ошибка")
}
