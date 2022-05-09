package usecase

import "csidealer/internal/entity"

type (
	Csi interface {
		HandleRawTraffic(data []byte)
		FlushBuffer()
		StartLog(filepath string)
		StopLog()
	}

	PackageRepo interface {
		Push(csiPackage *entity.Package)
		GetLastN(n int) []*entity.Package
		GetFullCount() uint64
	}

	RawTrafficRepo interface {
		Push(data []byte)
		GetAllSplitted() [][]byte
		Flush()
	}

	FileWriter interface {
		Start(filename string) error
		Stop()
		Write(data []byte) error
		IsOpen() bool
	}
)
