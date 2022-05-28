package usecase

import (
	"csidealer/internal/usecase/decoder"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CsiUseCase struct {
	repo             PackageRepo
	rawRepo          RawTrafficRepo
	fw               FileLogger
	csiPackageNumber uint64
}

func NewCsiUseCase(r PackageRepo, rR RawTrafficRepo, fw FileLogger) *CsiUseCase {
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
		uc.push(d.Data)
		uc.log(d.Data)
	}

	packets := uc.repo.GetLastN(1)
	if len(packets) > 0 {
		info := packets[0].Info
		fmt.Println(info)
	}
}

func (uc *CsiUseCase) StartLog(filepath string) error {
	err := uc.fw.Start(filepath)
	return err
}

func (uc *CsiUseCase) StopLog() {
	uc.fw.Stop()
}

func (uc *CsiUseCase) FlushBuffer() {
	fmt.Println("Буфер очищен!")
	uc.rawRepo.Flush()
}

func (uc *CsiUseCase) push(d []byte) {
	pack := decoder.DecodeCsiPackage(d)

	if pack.Info.CsiLength == 0 {
		return
	}

	pack.Uuid = uuid.New().String() // TODO: ? не совсем правильно держать это в usecase
	pack.Timestamp = time.Now().UnixMilli()
	pack.Number = uc.csiPackageNumber
	uc.csiPackageNumber += 1

	uc.repo.Push(pack)
}

func (uc *CsiUseCase) log(d []byte) {
	if !uc.fw.IsOpen() {
		return
	}

	uc.fw.Write(d)
}
