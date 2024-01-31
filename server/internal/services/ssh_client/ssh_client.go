package ssh_client

import (
	"errors"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/ssh"
)

type SshClient struct {
	isClientMainActive bool
	isSendData         bool
	isAvailable        bool
	uuid               string
	addr               string
	username           string
	conn               *ssh.Client
}

func NewSshClient(username string) *SshClient {
	return &SshClient{
		username: username,
		uuid:     uuid.New().String(),
	}
}

func (c *SshClient) GetId() string {
	return c.uuid
}

func (c *SshClient) command(text string) error {
	session, err := c.conn.NewSession()
	if err != nil {
		return err
	}

	if err := session.Run(text); err != nil {
		return err
	}

	return nil
}

func (c *SshClient) Connect(addr string) error {
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

func (c *SshClient) GetIsConnected() bool {
	return c.conn != nil
}

func (c *SshClient) Disconnect() error {
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

func (c *SshClient) GetAddr() string {
	return c.addr
}

func (c SshClient) ClientMainRun(serverIP string, serverPort int) error {
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

	log.Print("КОМАНДА-НАЧАТА")

	go c.command("client_main " + serverIP + " " + strconv.Itoa(serverPort))

	// if err := c.command("client_main " + serverIP + " " + serverPort); err != nil {
	// 	return err
	// }

	log.Print("КОМАНДА-ВЫПОЛНЕНА")

	return nil
}

func (c *SshClient) GetIsClientMainActive() bool {
	return c.isClientMainActive
}

func (c *SshClient) ClientMainStop() error {
	if err := c.command("killall client_main"); err != nil {
		return err
	}
	c.isClientMainActive = false
	return nil
}

func (c *SshClient) SendDataRun(ifName, dstMacAddr string, numOfPacketToSend, pktIntervalUs, pktLen uint16) error {
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

	go func() {
		log.Print("горутина sendData начата", c.isSendData)
		for c.isSendData {
			log.Print("sendData in cycle")
			c.command("sendData " +
				ifName + " " +
				dstMacAddr + " " +
				strconv.Itoa(int(numOfPacketToSend)) + " " +
				strconv.Itoa(int(pktIntervalUs)) + " " +
				strconv.Itoa(int(pktLen)))
		}
	}()

	// c.isSendData = false
	return nil
}

func (c *SshClient) GetIsSendDataActive() bool {
	return c.isSendData
}

func (c *SshClient) SendDataStop() error {
	if err := c.command("killall sendData"); err != nil {
		return err
	}
	c.isSendData = false
	return nil
}

func (c *SshClient) checkAvailable() {
	out, _ := exec.Command("ping", c.addr, "-c 5", "-i 3", "-w 10").Output()
	c.isAvailable = !strings.Contains(string(out), "Destination Host Unreachable")
}

func (c *SshClient) GetIsAvailable() bool {
	return c.isAvailable
}