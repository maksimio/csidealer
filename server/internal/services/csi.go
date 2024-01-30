package services

import entity "csidealer/internal/models"

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
	isFilterActive   bool

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
