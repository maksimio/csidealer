package ssh

import (
	"errors"
	"fmt"
	"strconv"

	"golang.org/x/crypto/ssh"
)

type AtherosClient struct {
	isClientMainActive bool
	isSendData         bool
	addr               string
	username           string
	conn               *ssh.Client
}

func NewAtherosClient(username string) *AtherosClient {
	return &AtherosClient{
		username: username,
	}
}

func (c *AtherosClient) command(text string) error {
	session, err := c.conn.NewSession()
	if err != nil {
		return err
	}

	fmt.Println(text)

	if err := session.Run(text); err != nil {
		return err
	}

	return nil
}

func (c *AtherosClient) Connect(addr string) error {
	config := &ssh.ClientConfig{
		User:            c.username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", addr+":22", config)
	if err != nil {
		return err
	}

	c.conn = conn
	c.addr = addr

	return nil
}

func (c *AtherosClient) GetIsConnected() bool {
	return c.conn != nil
}

func (c *AtherosClient) Disconnect() error {
	if !c.GetIsConnected() {
		return errors.New("нет соединения. Нечего отключать")
	}
	err := c.conn.Close()
	if err != nil {
		return err
	}

	c.conn = nil
	c.isClientMainActive = false //TODO что будет, если оставить активным и отключиться
	return nil
}

func (c *AtherosClient) GetAddr() string {
	return c.addr
}

func (c *AtherosClient) ClientMainStart(serverIP string, serverPort string) error {
	if err := c.command("client_main " + serverIP + " " + serverPort); err != nil {
		return err
	}
	c.isClientMainActive = true
	return nil
}

func (c *AtherosClient) GetIsClientMainActive() bool {
	return c.isClientMainActive
}

func (c *AtherosClient) ClientMainStop() error {
	if err := c.command("killall client_main"); err != nil {
		return err
	}
	c.isClientMainActive = false
	return nil
}

func (c *AtherosClient) SendDataRun(ifName, dstMacAddr string, numOfPacketToSend, pktIntervalUs, pktLen uint16) error {
	if c.isSendData {
		return errors.New("данные уже передаются")
	}
	if !c.GetIsConnected() {
		return errors.New("нет подключения, операция SendDataRun невозможна")
	}
	if c.isClientMainActive {
		return errors.New("нельзя отправлять данные, когда запущен сервер")
	}
	c.isSendData = true

	for c.isSendData {
		fmt.Println("SEND COMMAND !!!")
		err := c.command("sendData " +
			ifName + " " +
			dstMacAddr + " " +
			strconv.Itoa(int(numOfPacketToSend)) + " " +
			strconv.Itoa(int(pktIntervalUs)) + " " +
			strconv.Itoa(int(pktLen)))
		if err != nil {
			return err
		}
	}

	c.isSendData = false
	return nil
}

func (c *AtherosClient) GetIsSendData() bool {
	return c.isSendData
}

func (c *AtherosClient) SendDataStop() error {
	if err := c.command("killall sendData"); err != nil {
		return err
	}
	c.isSendData = false
	return nil
}
