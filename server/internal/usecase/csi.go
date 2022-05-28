package usecase

type CsiUseCase struct {
	repo    Repo
	rawRepo Buffer
	fl      FSLogger
	proc    Processor
	filter  Filter
	decoder Decoder

	csiPackageNumber uint64
	TcpRemoteAddr    string
	isFilterActive   bool
}

func NewCsiUseCase(repo Repo, rawRepo Buffer, fl FSLogger, proc Processor, filter Filter, decoder Decoder) *CsiUseCase {
	return &CsiUseCase{
		repo:    repo,
		rawRepo: rawRepo,
		fl:      fl,
		proc:    proc,
		filter:  filter,
		decoder: decoder,
	}
}
