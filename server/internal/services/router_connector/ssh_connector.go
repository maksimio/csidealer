package router_connector

import (
	"csidealer/internal/models"
	"csidealer/internal/services/router_connector/router"
	"errors"
	"log"
	"time"
)

const RECONNECT_DELAY = time.Second * 15

type RouterConnectorService struct {
	tx         router.Router
	rx         router.Router
	ServerIp   string
	ServerPort int
	SendData   models.SendDataInfo
}

func NewRouterConnectorService(txInfo, rxInfo models.RouterConfigInfo, ServerIp string, ServerPort int,
	sendData models.SendDataInfo) *RouterConnectorService {
	connector := &RouterConnectorService{
		tx: *router.NewRouter(txInfo),
		rx: *router.NewRouter(rxInfo),

		ServerIp:   ServerIp,
		ServerPort: ServerPort,
		SendData:   sendData,
	}

	connector.tx.Reconnect()
	connector.rx.Reconnect()

	go func() { // проверка связи
		for {
			if !connector.tx.CheckAvailable() {
				connector.tx.Reconnect()
				log.Println("TX RECONNECT")
			}
			if !connector.rx.CheckAvailable() {
				connector.rx.Reconnect()
				log.Println("RX RECONNECT")
			}
			time.Sleep(RECONNECT_DELAY)
		}
	}()

	return connector
}

func (s *RouterConnectorService) Reconnect() error {
	err1 := s.rx.Reconnect()
	err2 := s.tx.Reconnect()
	if err1 != nil || err2 != nil {
		return errors.New("не удалось подключиться к одному из роутеров")
	}

	return nil
}

func (s *RouterConnectorService) Start() error {
	// автоматически выполняем подключение
	if !s.rx.IsConnected() || !s.tx.IsConnected() {
		if err := s.Reconnect(); err != nil {
			return err
		}
	}

	if err := s.rx.RunClientMain(s.ServerIp, s.ServerPort); err != nil {
		return err
	}

	// RunClientMain не ждет старта, а просто запускает отдельную горутину. Это может привести к тому, что SendData будет запущен раньше
	// поэтому возникнет ошибка отправки пакетов. Чтобы этого избежать, используется простая задержка
	time.Sleep(time.Second)

	return s.tx.RunSendData(
		s.SendData.IfName,
		s.SendData.DstMacAddr,
		s.SendData.NumOfPacketToSend,
		s.SendData.PktIntervalUs,
		s.SendData.PktLen,
	)
}

func (s *RouterConnectorService) Stop() error {
	// автоматически выполняем подключение
	if !s.rx.IsConnected() || !s.tx.IsConnected() {
		if err := s.Reconnect(); err != nil {
			return err
		}
	}

	if err := s.tx.StopSendData(); err != nil {
		return err
	}

	return s.rx.StopClientMain()
}

func (s *RouterConnectorService) Status() models.ApiRoutersStatus {
	rx := models.ApiRouterInfo{
		Id:                 s.rx.Uuid,
		Addr:               s.rx.IpAddr,
		IsConnected:        s.rx.IsConnected(),
		IsClientMainActive: s.rx.IsClientMainActive,
		IsSendDataActive:   s.rx.IsSendData,
	}
	tx := models.ApiRouterInfo{
		Id:                 s.tx.Uuid,
		Addr:               s.tx.IpAddr,
		IsConnected:        s.tx.IsConnected(),
		IsClientMainActive: s.tx.IsClientMainActive,
		IsSendDataActive:   s.tx.IsSendData,
	}

	return models.ApiRoutersStatus{
		Rx:         rx,
		Tx:         tx,
		SendData:   s.SendData,
		ServerIp:   s.ServerIp,
		ServerPort: s.ServerPort,
	}
}
