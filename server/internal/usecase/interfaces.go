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
		GetCsiPackageCount() uint64
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

	Processor interface {
		Abs(data []*entity.Package) []entity.ApiPackage
		Phase(data []*entity.Package) []entity.ApiPackage
		Re(data []*entity.Package) []entity.ApiPackage
		Im(data []*entity.Package) []entity.ApiPackage
		PhaseWithoutJumps(data []*entity.Package) []entity.ApiPackage
	}
)
