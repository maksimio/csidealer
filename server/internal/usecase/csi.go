package usecase

type CsiUseCase struct {
	repo    PackageRepo
	rawRepo RawTrafficRepo
	fl      FileLogger
	proc    Processor

	csiPackageNumber uint64
	TcpRemoteAddr    string
}

func NewCsiUseCase(repo PackageRepo, rawRepo RawTrafficRepo, fl FileLogger, proc Processor) *CsiUseCase {
	return &CsiUseCase{
		repo:    repo,
		rawRepo: rawRepo,
		fl:      fl,
		proc:    proc,
	}
}
