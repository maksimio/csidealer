package usecase

import (
	"csidealer/internal/entity"
	"csidealer/internal/usecase/decoder"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (uc *CsiUseCase) StartLog(filepath string) error {
	err := uc.fl.Start(filepath)
	return err
}

func (uc *CsiUseCase) StopLog() error {
	if !uc.fl.IsOpen() {
		return errors.New("сейчас запись в файл не происходит. Нечего останавливать")
	}

	uc.fl.Stop()
	return nil
}

func (uc *CsiUseCase) MoveRawTraffic(data []byte) {
	uc.rawRepo.Push(data)
	splittedData := uc.rawRepo.GetAllSplitted()

	for _, d := range splittedData {
		uc.push(d.Data)
		uc.log(d)
	}

	// Просто вывод для теста
	packets := uc.repo.GetLastN(1)
	if len(packets) > 0 {
		info := packets[0].Info
		fmt.Println(info)
	}
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

	pack.Uuid = uuid.New().String()
	pack.Timestamp = time.Now().UnixMilli()
	pack.Number = uc.csiPackageNumber
	uc.csiPackageNumber += 1

	uc.repo.Push(pack)
}

func (uc *CsiUseCase) log(d entity.RawPackage) {
	if !uc.fl.IsOpen() {
		return
	}

	uc.fl.Write(d.Data)
}
