package usecase

type CsiUseCase struct {
	repo    IRepo
	rawRepo IBuffer
	fl      IFSLogger
	proc    IProcessor
	filter  IFilter
	decoder IDecoder

	csiPackageNumber uint64
	TcpRemoteAddr    string
	isFilterActive   bool
}

func NewCsiUseCase(repo IRepo, rawRepo IBuffer, fl IFSLogger, proc IProcessor, filter IFilter, decoder IDecoder) *CsiUseCase {
	return &CsiUseCase{
		repo:    repo,
		rawRepo: rawRepo,
		fl:      fl,
		proc:    proc,
		filter:  filter,
		decoder: decoder,
	}
}
