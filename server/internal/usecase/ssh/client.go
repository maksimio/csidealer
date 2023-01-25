package ssh

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/ssh"
)

type AtherosClient struct {
	isClientMainActive bool
	isSendData         bool
	isAvailable        bool
	uuid               string
	addr               string
	username           string
	conn               *ssh.Client
}

func NewAtherosClient(username string) *AtherosClient {
	return &AtherosClient{
		username: username,
		uuid:     uuid.New().String(),
	}
}

func (c *AtherosClient) GetId() string {
	return c.uuid
}

func (c *AtherosClient) command(text string) error {
	session, err := c.conn.NewSession()
	if err != nil {
		return err
	}

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

	c.ClientMainStop()
	c.SendDataStop()

	return nil
}

func (c *AtherosClient) GetIsConnected() bool {
	return c.conn != nil
}

func (c *AtherosClient) Disconnect() error {
	if !c.GetIsConnected() {
		return errors.New("нет соединения. Нечего отключать")
	}

	c.ClientMainStop()
	c.SendDataStop()

	err := c.conn.Close()
	if err != nil {
		return err
	}

	c.conn = nil
	return nil
}

func (c *AtherosClient) GetAddr() string {
	return c.addr
}

func (c AtherosClient) ClientMainRun(serverIP string, serverPort string) error {
	if c.isClientMainActive {
		return errors.New("client_main уже запущен")
	}
	if !c.GetIsConnected() {
		return errors.New("нет подключения, операция ClientMainRun невозможна")
	}
	if c.isSendData {
		return errors.New("нельзя принимать данные, когда отправляются данные")
	}

	c.isClientMainActive = true
	if err := c.command("client_main " + serverIP + " " + serverPort); err != nil {
		return err
	}

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

func (c *AtherosClient) GetIsSendDataActive() bool {
	return c.isSendData
}

func (c *AtherosClient) SendDataStop() error {
	if err := c.command("killall sendData"); err != nil {
		return err
	}
	c.isSendData = false
	return nil
}

func (c *AtherosClient) checkAvailable() {
	out, _ := exec.Command("ping", c.addr, "-c 5", "-i 3", "-w 10").Output()
	c.isAvailable = !strings.Contains(string(out), "Destination Host Unreachable")
}

func (c *AtherosClient) GetIsAvailable() bool {
	return c.isAvailable
}
