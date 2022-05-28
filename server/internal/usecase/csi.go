package usecase

type CsiUseCase struct {
	repo    PackageRepo
	rawRepo RawTrafficRepo
	fl      FileLogger
	proc    Processor
	filter  Filter

	csiPackageNumber uint64
	TcpRemoteAddr    string
	isFilterActive   bool
}

func NewCsiUseCase(repo PackageRepo, rawRepo RawTrafficRepo, fl FileLogger, proc Processor, filter Filter) *CsiUseCase {
	return &CsiUseCase{
		repo:    repo,
		rawRepo: rawRepo,
		fl:      fl,
		proc:    proc,
		filter:  filter,
	}
}
