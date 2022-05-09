package usecase

import "csidealer/internal/entity"

type (
	PackageRepo interface {
		Push(csiPackage *entity.Package)
		GetLastN(n int) []*entity.Package
	}

	RawTrafficRepo interface {
		Push(data []byte)
		GetAllSplitted() [][]byte
	}

	FileWriter interface {
		Start() error
		Stop() error
		Write(data []byte) error
	}
)