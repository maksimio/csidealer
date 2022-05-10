package usecase

import (
	"csidealer/internal/usecase/decoder"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type CsiUseCase struct {
	repo             PackageRepo
	rawRepo          RawTrafficRepo
	fw               FileWriter
	csiPackageNumber uint64
}

func NewCsiUseCase(r PackageRepo, rR RawTrafficRepo, fw FileWriter) *CsiUseCase {
	return &CsiUseCase{
		repo:    r,
		rawRepo: rR,
		fw:      fw,
	}
}

func (uc *CsiUseCase) HandleRawTraffic(data []byte) {
	uc.rawRepo.Push(data)
	splittedData := uc.rawRepo.GetAllSplitted()

	for _, d := range splittedData {
		pack := decoder.DecodeCsiPackage(d)

		pack.Uuid = uuid.New().String() // TODO: не совсем правильно держать это в usecase
		pack.Timestamp = time.Now().UnixMilli()
		pack.Number = uc.csiPackageNumber
		uc.csiPackageNumber += 1

		if pack.Info.CsiLength > 0 {
			uc.repo.Push(pack)
		}
	}

	packets := uc.repo.GetLastN(1)
	if len(packets) > 0 {
		info := packets[0].Info
		fmt.Println(info)
	}
}

func (uc *CsiUseCase) StartLog(filepath string) {

}

func (uc *CsiUseCase) StopLog() {

}

func (uc *CsiUseCase) FlushBuffer() {
	fmt.Println("Буфер очищен!")
	uc.rawRepo.Flush()
}
