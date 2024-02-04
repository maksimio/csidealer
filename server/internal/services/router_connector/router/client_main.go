package router

import (
	"errors"
	"log"
	"strconv"
)

// запускает TCP-клиент, пересылающий принятые данные CSI
func (c Router) RunClientMain(serverIP string, serverPort int) error {
	if !c.IsConnected() {
		return errors.New("нет подключения, операция ClientMainRun невозможна")
	}
	if c.IsClientMainActive {
		return errors.New("client_main уже запущен")
	}
	if c.IsSendData {
		return errors.New("нельзя принимать данные, когда отправляются данные")
	}

	c.IsClientMainActive = true
	log.Println("client_main - ретрансляция CSI начата, IP роутера:", c.IpAddr)
	go c.command("client_main " + serverIP + " " + strconv.Itoa(serverPort))
	log.Println("client_main - ретрансляция CSI прервана, IP роутера:", c.IpAddr)

	return nil
}

func (c *Router) StopClientMain() error {
	if !c.IsConnected() {
		return errors.New("нет подключения, операция StopClientMain невозможна")
	}
	if err := c.command("killall client_main"); err != nil {
		return err
	}

	c.IsClientMainActive = false
	return nil
}
