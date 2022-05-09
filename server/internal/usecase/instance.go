package usecase

type CsiUseCase struct {
	repo    PackageRepo
	rawRepo RawTrafficRepo
	fw      FileWriter
}

func NewCsiUseCase(r PackageRepo, rR RawTrafficRepo, fw FileWriter) *CsiUseCase {
	return &CsiUseCase{
		repo:    r,
		rawRepo: rR,
		fw:      fw,
	}
}


