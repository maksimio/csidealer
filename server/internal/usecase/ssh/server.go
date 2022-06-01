package ssh

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
)

type SshServer struct {
}

func NewSshServer() *SshServer {
	config := &ssh.ClientConfig{
		User:            "root",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", "192.168.1.1:22", config)
	if err != nil {
		fmt.Println(err)
		fmt.Println("СОЕДИНЕНИЕ НЕЕЕЕЕЕЕЕЕЕЕ установлено")
	} else {
		fmt.Println("СОЕДИНЕНИЕ УСТАНОВЛЕНО")
		defer conn.Close()
	}
	fmt.Println(conn)

	session, err := conn.NewSession()
	if err != nil {
		fmt.Println("ОШИБКА", err)
	}

	// // configure terminal mode
	// modes := ssh.TerminalModes{
	// 	ssh.ECHO: 0, // supress echo

	// }
	// // run terminal session
	// if err := session.RequestPty("xterm", 50, 80, modes); err != nil {
	// 	fmt.Println("ОШИБКА1", err)
	// }
	// start remote shell
	// if err := session.Shell(); err != nil {
	// 	fmt.Println("ОШИБКА2", err)
	// }

	var buff bytes.Buffer
	session.Stdout = &buff
	if err := session.Run("df -h"); err != nil {
		fmt.Println("ОШИБКА3", err)
	}
	fmt.Println(buff.String())

	return &SshServer{}
}
