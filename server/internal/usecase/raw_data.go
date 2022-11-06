package usecase

import (
	"csidealer/internal/entity"
	"csidealer/internal/usecase/processor"
	"encoding/binary"
	"errors"
	"time"

	"github.com/google/uuid"
)

func (uc *CsiUseCase) GetTcpRemoteAddr() string {
	return uc.TcpRemoteAddr
}

func (uc *CsiUseCase) SetTcpRemoteAddr(addr string) {
	uc.TcpRemoteAddr = addr
}

func (uc *CsiUseCase) StartLog(filepath string) error {
	err := uc.fl.Start(filepath)
	if err != nil {
		uc.logPackageCount = 0
	}
	return err
}

func (uc *CsiUseCase) StopLog() error {
	if !uc.fl.IsOpen() {
		return errors.New("сейчас запись в файл не происходит. Нечего останавливать")
	}

	uc.fl.Stop()
	return nil
}

func (uc *CsiUseCase) IsLog() bool {
	return uc.fl.IsOpen()
}

func (uc *CsiUseCase) GetLogWriteByteCount() uint64 {
	return uc.fl.GetWriteByteCount()
}

func (uc *CsiUseCase) GetLogStartTime() int64 {
	return uc.fl.GetStartTime()
}

func (uc *CsiUseCase) GetLogPackageCount() uint64 {
	return uc.logPackageCount
}

func (uc *CsiUseCase) MoveRawTraffic(data []byte) {
	uc.rawRepo.Push(data)
	splittedData := uc.rawRepo.GetAllSplitted()

	for _, d := range splittedData {
		// fmt.Println(uc.csiPackageNumber)
		uc.push(d.Data)
		uc.log(d)
	}
}

func (uc *CsiUseCase) FlushBuffer() {
	uc.rawRepo.Flush()
}

func (uc *CsiUseCase) push(d []byte) {
	pack := uc.decoder.DecodeCsiPackage(d)

	if pack.Info.CsiLength == 0 {
		return
	}

	if uc.isFilterActive && !uc.filter.Check(pack.Info) {
		return
	}

	pack.Uuid = uuid.New().String()
	pack.Timestamp = time.Now().UnixMilli()
	pack.Number = uc.csiPackageNumber
	uc.csiPackageNumber += 1

	apiPack := entity.ApiPackage{
		Timestamp: pack.Timestamp,
		Id:        pack.Uuid,
		Info:      pack.Info,
		Number:    pack.Number,
		Data:      uc.proc.CsiMap(pack.Data, processor.AbsHandler),
	}

	uc.cbPushPacket(apiPack)

	uc.repo.Push(pack)
}

func (uc *CsiUseCase) log(pack entity.RawPackage) {
	if !uc.fl.IsOpen() {
		return
	}

	bufSize16 := make([]byte, 2)
	binary.BigEndian.PutUint16(bufSize16, pack.Size)
	uc.fl.Write(bufSize16)
	uc.fl.Write(pack.Data)
	uc.logPackageCount += 1
}

func (uc *CsiUseCase) OnPushPacket(cb func(entity.ApiPackage)) {
	uc.cbPushPacket = cb
}
