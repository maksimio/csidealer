package usecase

import (
	"csidealer/internal/entity"
	"csidealer/internal/usecase/processor"
	"encoding/binary"
	"errors"
	"fmt"
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
		return err
	}
	uc.logPackageCount = 0
	return nil
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

	fmt.Println("PUSH ", uc.csiPackageNumber)

	uc.repo.Push(pack)

	// Сглаживание
	N := 10 // Порядок сглаживания
	abs := uc.proc.CsiMap(pack.Data, processor.AbsHandler)

	if uc.csiPackageNumber > uint64(N) {
		prevs := uc.repo.GetLastN(N)

		for i := 0; i < N; i++ {
			prev_abs := uc.proc.CsiMap(prevs[i].Data, processor.AbsHandler)
			for j := 0; j < 4; j++ {
				for k := 0; k < 56; k++ {
					abs[j][k] += prev_abs[j][k]
				}
			}
		}

		for j := 0; j < 4; j++ {
			for k := 0; k < 56; k++ {
				abs[j][k] /= float64(N)
			}
		}
	}
	uc.repo.GetLastN(3)
	// Конец сглаживания

	apiPack := entity.ApiPackageAbsPhase{
		Timestamp: pack.Timestamp,
		Id:        pack.Uuid,
		Info:      pack.Info,
		Number:    pack.Number,
		Abs:       abs,
		Phase:     uc.proc.CsiMap(pack.Data, processor.PhaseHandler),
	}

	uc.cbPushPacket(apiPack)
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

func (uc *CsiUseCase) OnPushPacket(cb func(entity.ApiPackageAbsPhase)) {
	uc.cbPushPacket = cb
}
