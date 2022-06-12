package usecase

import (
	"csidealer/internal/entity"
	"errors"
)

func (uc *CsiUseCase) getRouter(id string) (*IAtherosClient, error) {
	for _, value := range uc.routers {
		if (*value).GetId() == id {
			return value, nil
		}
	}

	return nil, errors.New("не существует устройства с id " + id)
}

func (uc *CsiUseCase) RoutersInfo() []entity.ApiRouterInfo {
	info := []entity.ApiRouterInfo{}
	for _, value := range uc.routers {
		router := (*value)
		info = append(info, entity.ApiRouterInfo{
			Id:                 router.GetId(),
			Addr:               router.GetAddr(),
			IsConnected:        router.GetIsConnected(),
			IsClientMainActive: router.GetIsClientMainActive(),
			IsSendDataActive:   router.GetIsSendDataActive(),
			IsAvailable:        router.GetIsAvailable(),
		})
	}

	return info
}

func (uc *CsiUseCase) RouterConnect(id, addr string) error {
	router, err := uc.getRouter(id)
	if err != nil {
		return err
	}

	return (*router).Connect(addr)
}

func (uc *CsiUseCase) RouterDisconnect(id string) error {
	router, err := uc.getRouter(id)
	if err != nil {
		return err
	}

	return (*router).Disconnect()
}

func (uc *CsiUseCase) RouterSendDataRun(id, DstMacAddr string, NumOfPacketToSend, pktIntervalUs, pktLen uint16) {
	router, err := uc.getRouter(id)
	if err != nil {
		return
	}

	go (*router).SendDataRun("wlan0", DstMacAddr, NumOfPacketToSend, pktIntervalUs, pktLen)
}

func (uc *CsiUseCase) RouterSendDataStop(id string) error {
	router, err := uc.getRouter(id)
	if err != nil {
		return err
	}

	return (*router).SendDataStop()
}

func (uc *CsiUseCase) RouterClientMainRun(id, ip, port string) {
	router, err := uc.getRouter(id)
	if err != nil {
		return
	}

	go (*router).ClientMainRun(ip, port)
}

func (uc *CsiUseCase) RouterClientMainStop(id string) error {
	router, err := uc.getRouter(id)
	if err != nil {
		return err
	}

	return (*router).ClientMainStop()
}
