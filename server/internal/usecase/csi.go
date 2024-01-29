package usecase

import "csidealer/internal/entity"

type CsiUseCase struct {
	repo        IRepo
	rawRepo     IBuffer
	fl          IFSLogger
	proc        IProcessor
	filter      IFilter
	decoder     IDecoder
	routers     []*IAtherosClient
	smoothOrder int

	csiPackageNumber uint64
	TcpRemoteAddr    string
	isFilterActive   bool
	logPackageCount  uint64

	cbPushPacket func(entity.ApiPackageAbsPhase)
}

func NewCsiUseCase(
	repo IRepo,
	rawRepo IBuffer,
	fl IFSLogger,
	proc IProcessor,
	filter IFilter,
	decoder IDecoder,
	routers []*IAtherosClient,
	smoothOrder int,
) *CsiUseCase {
	return &CsiUseCase{
		repo:        repo,
		rawRepo:     rawRepo,
		fl:          fl,
		proc:        proc,
		filter:      filter,
		decoder:     decoder,
		routers:     routers,
		smoothOrder: smoothOrder,
	}
}
