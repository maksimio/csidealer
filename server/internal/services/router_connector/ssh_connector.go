package router_connector

import (
	"csidealer/internal/models"
	"csidealer/internal/services/router_connector/router"
)

type RouterConnectorService struct {
	tx         router.Router
	rx         router.Router
	ServerIp   string
	ServerPort int

	IfName            string
	DstMacAddr        string
	NumOfPacketToSend uint16
	PktIntervalUs     uint16
	PktLen            uint16
}

func NewRouterConnectorService(txInfo, rxInfo models.RouterInfo, ServerIp string, ServerPort int,
	IfName, DstMacAddr string, NumOfPacketToSend, PktIntervalUs, PktLen uint16) *RouterConnectorService {
	return &RouterConnectorService{
		tx:                *router.NewRouter(txInfo),
		rx:                *router.NewRouter(rxInfo),
		ServerIp:          ServerIp,
		ServerPort:        ServerPort,
		IfName:            IfName,
		DstMacAddr:        DstMacAddr,
		NumOfPacketToSend: NumOfPacketToSend,
		PktIntervalUs:     PktIntervalUs,
		PktLen:            PktLen,
	}
}

func (s *RouterConnectorService) Connect() error {
	if err := s.rx.Connect(); err != nil {
		return err
	}

	return s.tx.Connect()
}

func (s *RouterConnectorService) Start() error {
	if err := s.rx.RunClientMain(s.ServerIp, s.ServerPort); err != nil {
		return err
	}

	return s.tx.RunSendData(
		s.IfName,
		s.DstMacAddr,
		s.NumOfPacketToSend,
		s.PktIntervalUs,
		s.PktLen,
	)
}

func (s *RouterConnectorService) Stop() error {
	if err := s.tx.StopSendData(); err != nil {
		return err
	}

	return s.rx.StopClientMain()
}
