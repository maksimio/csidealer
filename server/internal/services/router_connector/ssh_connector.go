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
	SendData   models.SendDataInfo
}

func NewRouterConnectorService(txInfo, rxInfo models.RouterInfo, ServerIp string, ServerPort int,
	sendData models.SendDataInfo) *RouterConnectorService {
	return &RouterConnectorService{
		tx: *router.NewRouter(txInfo),
		rx: *router.NewRouter(rxInfo),

		ServerIp:   ServerIp,
		ServerPort: ServerPort,
		SendData:   sendData,
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
		s.SendData.IfName,
		s.SendData.DstMacAddr,
		s.SendData.NumOfPacketToSend,
		s.SendData.PktIntervalUs,
		s.SendData.PktLen,
	)
}

func (s *RouterConnectorService) Stop() error {
	if err := s.tx.StopSendData(); err != nil {
		return err
	}

	return s.rx.StopClientMain()
}
