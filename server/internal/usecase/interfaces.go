package usecase

import "csidealer/internal/entity"

type (
	CsiUC interface {
		MoveRawTraffic(data []byte)
		FlushBuffer()
		StartLog(filepath string) error
		StopLog() error
		IsLog() bool
		GetTcpRemoteAddr() string
		SetTcpRemoteAddr(addr string)
		GetCsi(csiType uint8, count int) ([]entity.ApiPackage, error)
		GetSubcarrier(csiType uint8, count, h, i int) ([]float64, error)
		GetCsiPackageCount() uint64
		GetCsiPackageMaxCount() uint64
	}

	PackageRepo interface {
		Push(csiPackage *entity.Package)
		GetLastN(n int) []*entity.Package
		GetFullCount() uint64
		GetMaxCount() uint64
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

	Processor interface {
		PackageMap(data []*entity.Package, handler func(complex128) float64) []entity.ApiPackage
		SubcarrierMap(data []*entity.Package, handler func(complex128) float64, h, i int) ([]float64, error)
	}
)
