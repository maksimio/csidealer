package usecase

type CsiUseCase struct {
	repo             PackageRepo
	rawRepo          RawTrafficRepo
	fl               FileLogger
	csiPackageNumber uint64
	TcpRemoteAddr    string
}

func NewCsiUseCase(repo PackageRepo, rawRepo RawTrafficRepo, fl FileLogger) *CsiUseCase {
	return &CsiUseCase{
		repo:    repo,
		rawRepo: rawRepo,
		fl:      fl,
	}
}
