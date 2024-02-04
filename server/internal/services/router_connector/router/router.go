package router

import (
	"csidealer/internal/models"
	"errors"
	"os/exec"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/ssh"
)

type Router struct {
	IsClientMainActive bool
	IsSendData         bool
	IsAvailable        bool

	Uuid     string
	IpAddr   string
	Username string

	conn *ssh.Client
}

func NewRouter(info models.RouterInfo) *Router {
	return &Router{
		Username: info.Username,
		IpAddr:   info.IpAddr,
		Uuid:     uuid.New().String(),
	}
}

func (c *Router) command(text string) error {
	session, err := c.conn.NewSession()
	if err != nil {
		return err
	}

	return session.Run(text)
}

func (c *Router) Reconnect() error {
	config := &ssh.ClientConfig{
		User:            c.Username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 5,
	}

	conn, err := ssh.Dial("tcp", c.IpAddr+":22", config)
	if err != nil {
		return err
	}

	c.conn = conn
	return nil
}

func (c *Router) IsConnected() bool {
	return c.conn != nil
}

func (c *Router) Disconnect() error {
	if !c.IsConnected() {
		return errors.New("нет соединения. Нечего отключать")
	}

	err := c.conn.Close()
	if err != nil {
		return err
	}

	c.conn = nil
	return nil
}

func (c *Router) checkAvailable() {
	out, _ := exec.Command("ping", c.IpAddr, "-c 5", "-i 3", "-w 10").Output()
	c.IsAvailable = !strings.Contains(string(out), "Destination Host Unreachable")
}
