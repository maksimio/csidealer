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

const SSH_CONNECT_TIMEOUT = time.Second * 5

type Router struct {
	IsClientMainActive bool
	IsSendData         bool

	Uuid     string
	IpAddr   string
	Username string

	conn *ssh.Client
}

func NewRouter(info models.RouterConfigInfo) *Router {
	router := &Router{
		Username: info.Username,
		IpAddr:   info.IpAddr,
		Uuid:     uuid.New().String(),
	}

	return router
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
		Timeout:         SSH_CONNECT_TIMEOUT,
	}

	if c.conn != nil {
		c.conn.Close()
	}

	conn, err := ssh.Dial("tcp", c.IpAddr+":22", config)
	c.conn = conn
	if err != nil {
		return err
	}

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

func (c *Router) CheckAvailable() bool {
	out, _ := exec.Command("ping", c.IpAddr, "-c 5", "-i 3", "-w 10").Output()
	isAvailable := !strings.Contains(string(out), "100% packet loss")
	return isAvailable
}
