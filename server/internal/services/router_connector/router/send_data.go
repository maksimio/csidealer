package router

import (
	"errors"
	"log"
	"strconv"
)

// запускает передачу данных с флагом CSI на роутер, на котором запущен client_main
func (c *Router) RunSendData(ifName, dstMacAddr string, numOfPacketToSend, pktIntervalUs, pktLen uint16) error {
	if !c.isConnected() {
		return errors.New("нет подключения, операция SendDataRun невозможна")
	}
	if c.IsSendData {
		return errors.New("данные уже передаются")
	}
	if c.IsClientMainActive {
		return errors.New("нельзя отправлять данные, когда запущен client_main")
	}
	c.IsSendData = true

	go func() {
		for c.IsSendData {
			log.Println("запуск sendData, IP:", c.IpAddr)
			c.command("sendData " +
				ifName + " " +
				dstMacAddr + " " +
				strconv.Itoa(int(numOfPacketToSend)) + " " +
				strconv.Itoa(int(pktIntervalUs)) + " " +
				strconv.Itoa(int(pktLen)))
		}
	}()

	return nil
}

func (c *Router) StopSendData() error {
	if !c.isConnected() {
		return errors.New("нет подключения, операция StopSendData невозможна")
	}
	c.IsSendData = false // сначала нужно убрать флаг, иначе после killall запустится следующая итерация sendData
	if err := c.command("killall sendData"); err != nil {
		return err
	}
	return nil
}
