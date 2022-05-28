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
		GetPackageFilterLimits() (isActive bool, payloadLenMin, payloadLenMax uint16, nr, nc, nTones uint8)
		SetPackageFilterLimits(isActive bool, payloadLenMin, payloadLenMax uint16, nr, nc, nTones uint8)
	}

	Repo interface {
		Push(csiPackage *entity.Package)
		GetLastN(n int) []*entity.Package
		GetFullCount() uint64
		GetMaxCount() uint64
	}

	Buffer interface {
		Push(data []byte)
		GetAllSplitted() []entity.RawPackage
		Flush()
	}

	FSLogger interface {
		Start(filename string) error
		Stop()
		Write(data []byte) error
		IsOpen() bool
	}

	Processor interface {
		PackageMap(data []*entity.Package, handler func(complex128) float64) []entity.ApiPackage
		SubcarrierMap(data []*entity.Package, handler func(complex128) float64, h, i int) ([]float64, error)
	}

	Filter interface {
		Check(info *entity.PackageInfo) bool
		GetLimits() (payloadLenMin, payloadLenMax uint16, nr, nc, nTones uint8)
		SetLimits(payloadLenMin, payloadLenMax uint16, nr, nc, nTones uint8)
	}

	Decoder interface {
		DecodeCsiPackage([]byte) *entity.Package
	}
)
