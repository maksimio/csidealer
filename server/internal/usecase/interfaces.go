package usecase

import "csidealer/internal/entity"

type (
	Csi interface {
		HandleRawTraffic(data []byte)
		FlushBuffer()
		StartLog(filepath string) error
		StopLog()
	}

	PackageRepo interface {
		Push(csiPackage *entity.Package)
		GetLastN(n int) []*entity.Package
		GetFullCount() uint64
	}

	RawTrafficRepo interface {
		Push(data []byte)
		GetAllSplitted() []entity.RawPackage
		Flush()
	}

	FileLogger interface {
		Start(filename string) error
		Stop()
		Write(data []byte) error
		IsOpen() bool
	}
)
